package handler

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) tokenVerification(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get(authorizationHeader)
	if authHeader == "" {
		respondError(w, http.StatusBadRequest, errors.New("auth header is empty"))
		return
	}

	splitToken := strings.Split(authHeader, " ")
	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		respondError(w, http.StatusBadRequest, errors.New("invalid auth type"))
		return
	}

	userID, err := h.svc.User.ParseToken(splitToken[1])
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	ctx := context.WithValue(r.Context(), "userID", userID)
	*r = *r.WithContext(ctx)
}

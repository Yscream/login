package handler

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func respondJSON(w http.ResponseWriter, code int, response any) {
	b, err := json.Marshal(response)
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(b)
}

func respondError(w http.ResponseWriter, code int, err error) {
	respondJSON(w, code, &Error{
		Code:    code,
		Message: err.Error(),
	})
}

package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Yscream/login/pkg/models"
	"github.com/Yscream/login/pkg/service"
)

type Handler struct {
	svc *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	var user *models.User
	err = json.Unmarshal(b, &user)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	token, err := h.svc.User.GenerateToken(user.Username, user.Password)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	respondJSON(w, http.StatusOK, token)
}

func (h *Handler) UploadPicture(w http.ResponseWriter, r *http.Request) {
	h.tokenVerification(w, r)
	userID := r.Context().Value("userID").(uint)

	r.ParseMultipartForm(10 * 1024 * 1024)

	file, handler, err := r.FormFile("picture")
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}
	defer file.Close()

	basePath, err := os.Getwd()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	folder := "pictures"

	folderPath := filepath.Join(basePath, folder)

	err = os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	filePath := filepath.Join(folderPath, handler.Filename)

	outputFile, err := os.Create(filePath)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, file)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	imagePath := fmt.Sprintf("%s/%s", folderPath, handler.Filename)

	imageURL := fmt.Sprintf("https://github.com/Yscream/login/%s/%s", folder, handler.Filename)

	image := &models.Image{
		UserID:    userID,
		ImagePath: imagePath,
		ImageURL:  imageURL,
	}

	err = h.svc.Image.SaveImage(image)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	respondJSON(w, http.StatusOK, "File uploaded successfully")
}

func (h *Handler) FetchImages(w http.ResponseWriter, r *http.Request) {
	//i know about 2 the most common ways to send image:
	//1) send path where image is
	//2) send image as base64 format
	//here i decided to use first one

	imagePathes, err := h.svc.Image.FetchImagePathes()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	respondJSON(w, http.StatusOK, imagePathes)
}

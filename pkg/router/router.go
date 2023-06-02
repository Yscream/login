package router

import (
	"net/http"

	"github.com/Yscream/login/pkg/handler"
	"github.com/Yscream/login/pkg/service"
	"github.com/gorilla/mux"
)

func NewRouter(svc *service.Service) *mux.Router {
	h := handler.NewHandler(svc)

	r := mux.NewRouter()

	r.HandleFunc("/login", h.Login).Methods(http.MethodPost)
	r.HandleFunc("/upload_picture", h.UploadPicture).Methods(http.MethodPost)
	r.HandleFunc("/images", h.FetchImages).Methods(http.MethodGet)

	return r
}

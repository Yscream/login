package service

import (
	"github.com/Yscream/login/pkg/models"
	"github.com/Yscream/login/pkg/repository"
)

type User interface {
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (uint, error)
}

type Image interface {
	SaveImage(image *models.Image) error
	FetchImagePathes() ([]string, error)
}

type Service struct {
	User
	Image
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:  NewUserService(repos.User),
		Image: NewImageService(repos.Image),
	}
}

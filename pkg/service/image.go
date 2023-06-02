package service

import (
	"github.com/Yscream/login/pkg/models"
	"github.com/Yscream/login/pkg/repository"
)

type ImageService struct {
	repo repository.Image
}

func NewImageService(repo repository.Image) *ImageService {
	return &ImageService{
		repo: repo,
	}
}

func (svc *ImageService) SaveImage(image *models.Image) error {
	err := svc.repo.SaveImage(image)
	if err != nil {
		return err
	}

	return nil
}

func (svc *ImageService) FetchImagePathes() ([]string, error) {
	imagePathes, err := svc.repo.FetchImagePathes()
	if err != nil {
		return nil, err
	}

	return imagePathes, nil
}
package repository

import (
	"github.com/Yscream/login/pkg/models"
	"github.com/Yscream/login/pkg/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type User interface {
	GetUser(username, passwordHash string) (*models.User, error)
}

type Image interface {
	SaveImage(image *models.Image) error
	FetchImagePathes() ([]string, error)
}

type Repository struct {
	User
	Image
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: postgres.NewUserRepository(db),
		Image: postgres.NewImageRepository(db),
	}
}

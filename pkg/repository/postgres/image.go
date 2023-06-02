package postgres

import (
	"fmt"

	"github.com/Yscream/login/pkg/models"
	"github.com/jmoiron/sqlx"
)

type ImageRepository struct {
	db *sqlx.DB
}

func NewImageRepository(db *sqlx.DB) *ImageRepository {
	return &ImageRepository{
		db: db,
	}
}

func (repo *ImageRepository) SaveImage(image *models.Image) error {
	query := "INSERT INTO images(user_id, image_path, image_url) VALUES($1, $2, $3)"

	_, err := repo.db.Exec(query, image.UserID, image.ImagePath, image.ImageURL)
	if err != nil {
		return fmt.Errorf("couldn't insert image")
	}
	
	return nil
}

func (repo* ImageRepository) FetchImagePathes() ([]string, error) {
	query := "SELECT image_path FROM images"

	images := make([]string, 0)
	err := repo.db.Select(&images, query)
	if err != nil {
		return nil, fmt.Errorf("couldn't select image pathes, %w", err)
	}

	return images, nil
}
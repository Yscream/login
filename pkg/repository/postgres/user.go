package postgres

import (
	"fmt"

	"github.com/Yscream/login/pkg/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) GetUser(username, passwordHash string) (*models.User, error) {
	fmt.Println(username, passwordHash)
	query := "SELECT * FROM users WHERE username=$1 AND password_hash=$2"

	user := &models.User{}
	err := repo.db.Get(user, query, username, passwordHash)
	if err != nil {
		return nil, fmt.Errorf("couldn't get user, %w", err)
	}

	return user, nil
}

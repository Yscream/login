package service

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/Yscream/login/pkg/repository"
	"github.com/golang-jwt/jwt"
)

const (
	tokenTTL   = 12 * time.Hour
	signingKey = "shcjks123;$5ll231#$%7563dvbz4%21^21$213"
)

type TokenClaims struct {
	UserID uint
	jwt.StandardClaims
}

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) GenerateToken(username, password string) (string, error) {
	user, err := svc.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	fmt.Println(user.ID)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		},
	})

	return token.SignedString([]byte(signingKey))
}

func (svc *UserService) ParseToken(accessToken string) (uint, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, fmt.Errorf("couldn't parse token: %w", err)
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, fmt.Errorf("invalid token")
	}

	return claims.UserID, nil
}

func generatePasswordHash(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(nil)))
}

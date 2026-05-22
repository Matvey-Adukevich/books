package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/Matvey-Adukevich/books"
	"github.com/Matvey-Adukevich/books/pkg/repository"
)

const salt = "lhj5kds71ionh30dsh763jgf5o8"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user books.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

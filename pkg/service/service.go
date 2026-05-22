package service

import (
	"github.com/Matvey-Adukevich/books"
	"github.com/Matvey-Adukevich/books/pkg/repository"
)

type Authorization interface {
	CreateUser(user books.User) (int, error)
}

type Book interface{}

type Service struct {
	Authorization
	Book
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}

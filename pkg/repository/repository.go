package repository

import (
	"github.com/Matvey-Adukevich/books"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user books.User) (int, error)
}

type Book interface{}

type Repository struct {
	Authorization
	Book
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}

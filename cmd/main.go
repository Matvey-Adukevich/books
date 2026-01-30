package main

import (
	"log"

	"github.com/Matvey-Adukevich/books"
	"github.com/Matvey-Adukevich/books/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	server := new(books.Server)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

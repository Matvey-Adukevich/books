package main

// import (
// 	// "database/sql"
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"os"

// 	// "context"

// 	"github.com/jmoiron/sqlx"
// 	"github.com/joho/godotenv"

// 	// "github.com/lib/pq"
// 	_ "github.com/jackc/pgx/v5/stdlib"
// )

// type Book struct {
// 	ID          string `json:"id"`
// 	Title       string `json:"title"`
// 	Author      string `json:"author"`
// 	Year        int    `json:"year"`
// 	Description string `json:"description"`
// }

// var db *sqlx.DB

// func main() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal("failed to load .env")
// 	}

// 	// var err error
// 	// connectionString := "postgres://postgres:123@localhost:5432/library?sslmode=disable"
// 	// connectionString := "host=localhost port=5432 user=postgres password=123 dbname=library sslmode=disable"
// 	// connectionString := "postgres://postgres@localhost:5432/library?sslmode=disable"
// 	// db, err := sql.Open("postgres", connectionString)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// defer db.Close()

// 	// if err = db.Ping(); err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// var err error
// 	// connectionString := "postgres://postgres:1234@localhost:5432/library?sslmode=disable"
// 	// connectionString := "postgres://postgres:1234@localhost:5432/library?sslmode=disable"

// 	// connectionString := "postgres://postgres@localhost:5432/library?sslmode=disable"
// 	log.Println("Connecting to bd")
// 	var err error
// 	db, err = sqlx.Connect("pgx", os.Getenv("DATABASE_URL"))
// 	if err != nil {
// 		log.Fatal("failed connect to db", err)
// 	}
// 	defer db.Close()
// 	// db, err = sql.Open("postgres", connectionString)
// 	// if err != nil {
// 	// 	log.Fatal("Ошибка Open:", err)
// 	// }
// 	// defer db.Close()

// 	// log.Println("Check Ping...")
// 	// if err = db.Ping(); err != nil {
// 	// 	log.Fatal("Failed Ping: ", err)
// 	// }

// 	log.Println("Connected to db")

// 	http.HandleFunc("/books", getBookList)
// 	// http.HandleFunc("/books/insert", insertBookInGlobalList)
// 	// http.HandleFunc("/books/", getBookByID)

// 	log.Println("Server :8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// // func insertBookInGlobalList(w http.ResponseWriter, r *http.Request){

// // }

// func getBookList(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var books []Book
// 	err := db.Select(&books, "SELECT id, title, author, year, description FROM books ORDER BY title")
// 	// rows, err := conn.Query("SELECT id, title, author, year, description FROM books ORDER BY title")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	// defer rows.Close()

// 	// var books []Book
// 	// for rows.Next() {
// 	// 	var b Book
// 	// 	err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Year, &b.Description)
// 	// 	if err != nil {
// 	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	// 		return
// 	// 	}
// 	// 	books = append(books, b)
// 	// }
// 	w.Header().Set("Content-Type", "application/json; charset=utf8")
// 	json.NewEncoder(w).Encode(books)
// }

import (
	"fmt"
	"log"
	"os"

	"github.com/Matvey-Adukevich/books"
	"github.com/Matvey-Adukevich/books/pkg/handler"
	"github.com/Matvey-Adukevich/books/pkg/repository"
	"github.com/Matvey-Adukevich/books/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DATABASE_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	var version string
	db.QueryRow("SELECT version()").Scan(&version)
	fmt.Println("ПОДКЛЮЧЕНО К:", version)

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(books.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

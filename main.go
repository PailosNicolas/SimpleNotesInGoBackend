package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/PailosNicolas/SimpleNotesInGoBackend/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	DB        *database.Queries
	jwtSecret string
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("CONN")
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		println("error starting DB")
		return
	}

	config := apiConfig{
		DB:        database.New(db),
		jwtSecret: os.Getenv("JWT_SECRET"),
	}

	port := os.Getenv("PORT")

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{}))

	// Users
	r.Post("/users/", config.HandlerCreateNewUser)
	r.Post("/login/", config.HandlerLogin)

	// Notes
	r.Post("/notes/", config.middlewareAuth(config.HandlerCreateNote))
	r.Put("/notes/", config.middlewareAuth(config.HandlerUpdateNote))
	r.Get("/notes/", config.middlewareAuth(config.HandlerGetNote))

	// Categories
	r.Post("/categories/", config.middlewareAuth(config.HandlerCreateCategory))
	r.Get("/categories/", config.middlewareAuth(config.HandlerGetCategory))
	r.Delete("/categories/", config.middlewareAuth(config.HandlerDeleteCategory))
	r.Post("/categories/notes/", config.middlewareAuth(config.HandlerAssingCategoryToNote))

	println("Server started on port: " + string(port))
	log.Fatal(http.ListenAndServe(":"+port, r))
}

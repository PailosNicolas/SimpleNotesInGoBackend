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
	DB *database.Queries
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
		DB: database.New(db),
	}

	port := os.Getenv("PORT")

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{}))

	// Users
	r.Post("/users/", config.HandlerCreateNewUser)

	log.Fatal(http.ListenAndServe(":"+port, r))
}

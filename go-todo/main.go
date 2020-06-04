package main

import (
	"fmt"
	"go-todo/domain"
	"go-todo/handlers"
	"go-todo/postgres"
	"log"
	"net/http"
	"os"

	"github.com/go-pg/pg/v9"
)

func main() {
	DB := postgres.New(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "go-todo",
	})
	defer DB.Close()

	domainDB := domain.DB{UserRepo: postgres.NewUserRepo(DB)}

	d := &domain.Domain{DB: domainDB}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	handler := handlers.SetupRouter(d)

	log.Printf("Listening on %s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

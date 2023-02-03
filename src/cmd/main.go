package main

import (
	"audiience_challenge/entrypoints"
	"audiience_challenge/repositories/rates"
	"audiience_challenge/services"
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {

	// DB initialization
	db, err := sql.Open("sqlite3", "./resources/rates.db")

	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sql.DB) {
		if err = db.Close(); err != nil {
			log.Fatal(err)
		}
	}(db)

	// Dependency injection starts
	repo := rates.NewRepository(db)
	estimateService := services.NewEstimateService(repo)
	router := mux.NewRouter().StrictSlash(true)
	server := entrypoints.NewServer(estimateService, router)
	server.SetupRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}

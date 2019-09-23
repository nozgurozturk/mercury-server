package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nozgurozturk/startpage_server/handlers"
	"github.com/nozgurozturk/startpage_server/models"
)

var pages []models.Page

func main() {

	// Create Router
	models.InitMigration()
	r := mux.NewRouter()
	// REST
	r.HandleFunc("/api/pages", handlers.GetPages).Methods("GET")
	r.HandleFunc("/api/pages", handlers.CreatePage).Methods("POST")
	r.HandleFunc("/api/page/{id}", handlers.DeletePage).Methods("DELETE")
	r.HandleFunc("/api/page/{id}", handlers.UpdatePage).Methods("PUT")
	r.HandleFunc("/api/page/{id}", handlers.GetPage).Methods("GET")
	// Server is started on localhost:8000
	fmt.Printf("Server is 🚀 on 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

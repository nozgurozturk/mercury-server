package main

import (
	"fmt"
	"log"
	"handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create Router
	r := mux.NewRouter()
	// REST
	r.HandleFunc("/api/pages", handlers.GetPages).Methods("GET")
	r.HandleFunc("/api/pages", handlers.CreatePage).Methods("POST")
	r.HandleFunc("/api/pages/{id}", handlers.DeletePage).Methods("DELETE")
	r.HandleFunc("/api/pages/{id}", handlers.UpdatePage).Methods("PUT")
	r.HandleFunc("/api/pages/{id}", handlers.GetPage).Methods("GET")
	// Server is started on localhost:8000
	fmt.Printf("Server is 🚀  on 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

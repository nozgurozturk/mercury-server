package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Page is collection of usefull source
type Page struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

// Project is collection of office's developed website
type Project struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Link        string `json:"link"`
	TestLink    string `json:"testlink"`
}

// Guide is collection of office's rule and guide
type Guide struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

var pages []Page
var projects []Project
var guides []Guide

func main() {
	// Mock Data
	pages = append(pages, Page{ID: "01", Description: "Google", Link: "https://www.google.com"})
	projects = append(projects, Project{ID: "01", Description: "Go Lang", Link: "https://www.golang.org", TestLink: "https://www.golang.test.org"})
	guides = append(guides, Guide{ID: "01", Description: "Github", Link: "http://www.github.com"})
	// Create Router
	r := mux.NewRouter()
	// REST
	r.HandleFunc("/api/pages", getPages).Methods("GET")
	r.HandleFunc("/api/pages", createPage).Methods("POST")
	r.HandleFunc("/api/pages/{id}", deletePage).Methods("DELETE")
	r.HandleFunc("/api/pages/{id}", updatePage).Methods("PUT")
	r.HandleFunc("/api/pages/{id}", getPage).Methods("GET")
	// Server is started on localhost:8000
	fmt.Printf("Server is 🚀  on 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func createPage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var page Page
	_ = json.NewDecoder(r.Body).Decode(&page)
	page.ID = "02"
	pages = append(pages, page)
	json.NewEncoder(w).Encode(page)
}
func getPages(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(pages)
}
func getPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r)
	for _, page := range pages {
		if page.ID == params["id"] {
			json.NewEncoder(w).Encode(pages)
		}
	}
	json.NewEncoder(w).Encode(&Page{})
}
func updatePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r)
	for index, item := range pages {
		if item.ID == params["id"] {
			pages = append(pages[:index], pages[index+1:]...)
			var page Page
			_ = json.NewDecoder(r.Body).Decode(&page)
			page.ID = item.ID
			pages = append(pages, page)
			json.NewEncoder(w).Encode(page)
			return
		}
	}
	json.NewEncoder(w).Encode(pages)
}

func deletePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r)
	for index, item := range pages {
		if item.ID == params["id"] {
			pages = append(pages[:index], pages[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(pages)
}

package handlers

import (
"encoding/json"
"models"
"net/http"

"github.com/gorilla/mux"
)

var pages []models.Page

var CreatePage = func(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var page models.Page
	_ = json.NewDecoder(r.Body).Decode(&page)
	page.ID = "02"
	pages = append(pages, page)
	json.NewEncoder(w).Encode(page)
}
var GetPages= func(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(pages)
}
var GetPage = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r)
	for _, page := range pages {
		if page.ID == params["id"] {
			json.NewEncoder(w).Encode(pages)
		}
	}
	json.NewEncoder(w).Encode(&models.Page{})
}
var UpdatePage = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r)
	for index, item := range pages {
		if item.ID == params["id"] {
			pages = append(pages[:index], pages[index+1:]...)
			var page models.Page
			_ = json.NewDecoder(r.Body).Decode(&page)
			page.ID = item.ID
			pages = append(pages, page)
			json.NewEncoder(w).Encode(page)
			return
		}
	}
	json.NewEncoder(w).Encode(pages)
}

var DeletePage= func(w http.ResponseWriter, r *http.Request) {
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

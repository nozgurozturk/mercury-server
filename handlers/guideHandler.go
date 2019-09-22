package handlers

import (
	"encoding/json"
	"net/http"

	"models"

	"github.com/gorilla/mux"
)

var guides []models.Guide

var CreateGuide = func(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var guide models.Guide
	_ = json.NewDecoder(r.Body).Decode(&guide)
	guide.ID = "02"
	guides = append(guides, guide)
	json.NewEncoder(w).Encode(guide)
}
var GetGuides = func(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(guides)
}
var GetGuide = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r)
	for _, guide := range guides {
		if guide.ID == params["id"] {
			json.NewEncoder(w).Encode(guides)
		}
	}
	json.NewEncoder(w).Encode(&models.Guide{})
}
var UpdateGuide = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r)
	for index, item := range guides {
		if item.ID == params["id"] {
			guides = append(guides[:index], guides[index+1:]...)
			var guide models.Guide
			_ = json.NewDecoder(r.Body).Decode(&guide)
			guide.ID = item.ID
			guides = append(guides, guide)
			json.NewEncoder(w).Encode(guide)
			return
		}
	}
	json.NewEncoder(w).Encode(guides)
}

var DeleteGuide = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r)
	for index, item := range guides {
		if item.ID == params["id"] {
			guides = append(guides[:index], guides[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(guides)
}

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nozgurozturk/startpage_server/models"
)

var pages []models.Page

var CreatePage = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var page models.Page
	_ = json.NewDecoder(r.Body).Decode(&page)
	models.GetDB().NewRecord(pages)
	models.GetDB().Create(&page)
	json.NewEncoder(w).Encode(page)
}
var GetPages = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	models.GetDB().Find(&pages)
	json.NewEncoder(w).Encode(pages)
}

var GetPage = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r)
	models.GetDB().Table("pages").Where("id = ?", params["id"]).First(&pages)
	json.NewEncoder(w).Encode(pages)
}
var UpdatePage = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r)
	var page models.Page
	_ = json.NewDecoder(r.Body).Decode(&page)
	models.GetDB().Table("pages").Where("id = ?", params["id"]).Save(&page)
	json.NewEncoder(w).Encode(pages)
}

var DeletePage = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r)
	var page models.Page
	models.GetDB().Table("pages").Where("id = ?", params["id"]).Delete(&page)
	json.NewEncoder(w).Encode(pages)
}

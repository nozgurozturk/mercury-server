package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nozgurozturk/startpage_server/models"

	"github.com/gorilla/mux"
)

var projects []models.Project

var CreateProject = func(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var project models.Project
	_ = json.NewDecoder(r.Body).Decode(&project)
	projects = append(projects, project)
	json.NewEncoder(w).Encode(project)
}
var GetProjects = func(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(projects)
}
var GetProject = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r)
	for _, project := range projects {
		if project.ID == params["id"] {
			json.NewEncoder(w).Encode(projects)
		}
	}
	json.NewEncoder(w).Encode(&models.Project{})
}
var UpdateProject = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r)
	for index, item := range projects {
		if item.ID == params["id"] {
			projects = append(projects[:index], projects[index+1:]...)
			var project models.Project
			_ = json.NewDecoder(r.Body).Decode(&project)
			project.ID = item.ID
			projects = append(projects, project)
			json.NewEncoder(w).Encode(project)
			return
		}
	}
	json.NewEncoder(w).Encode(projects)
}

var DeleteProject = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	params := mux.Vars(r)
	for index, item := range projects {
		if item.ID == params["id"] {
			projects = append(projects[:index], projects[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(projects)
}

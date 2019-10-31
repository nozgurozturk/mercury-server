package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nozgurozturk/mercury-server/api/models"
	"github.com/nozgurozturk/mercury-server/api/utils"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (server *Server) CreateLink (w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		utils.ERROR(w , http.StatusUnprocessableEntity, err)
		return
	}
	link := models.Link{}
	err = json.Unmarshal(body, &link)
	if err != nil{
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = link.Validate()
	if err != nil{
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	linkCreated , err := link.SaveLink(server.DB)
	if err != nil {
		formattedError := utils.ErrorType(err.Error())
		utils.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, linkCreated.ID))
	utils.Respond(w, http.StatusOK, linkCreated)
}

func (server *Server) GetLinks (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	iid,err := strconv.ParseInt(vars["id"], 10, 64)

	if  err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	link := models.Link{}
	links, err :=link.FindAllLink(server.DB, uint32(iid))
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	utils.Respond(w, http.StatusOK, links)
}

func (server *Server) GetLink (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	lid, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	link := models.Link{}
	selectedLink, err := link.FindLink(server.DB, uint32(lid))
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	utils.Respond(w, http.StatusOK, selectedLink)
}

func (server *Server) UpdateLink (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	lid, err := strconv.ParseInt(vars["id"], 10,64)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	link := models.Link{}
	err = server.DB.Model(models.Link{}).Where("id = ?", lid).First(&link).Error
	if err != nil {
		utils.ERROR(w, http.StatusNotFound, errors.New("link not found"))
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	linkUpdate := models.Link{}
	err = json.Unmarshal(body, &linkUpdate)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = linkUpdate.Validate()
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedLink, err := linkUpdate.UpdateLink(server.DB, uint32(lid))
	if err != nil {
		formattedError := utils.ErrorType(err.Error())
		utils.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	utils.Respond(w, http.StatusOK, updatedLink)
}

func (server *Server) DeleteLink (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	lid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	link := models.Link{}
	err = server.DB.Debug().Model(models.Link{}).Where("id = ?", lid).First(&link).Error
	if err != nil {
		utils.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	_, err = link.DeleteLink(server.DB, uint32(lid))
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", lid))
	utils.Respond(w, http.StatusNoContent, "")
}


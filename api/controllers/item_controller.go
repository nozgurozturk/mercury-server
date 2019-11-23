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

func (server *Server) CreateItem (w http.ResponseWriter, r *http.Request){

	body, err := ioutil.ReadAll(r.Body)

	if err != nil{
		utils.ERROR(w , http.StatusUnprocessableEntity, err)
		return
	}

	item := models.Item{}


	err = json.Unmarshal(body, &item)
	if err != nil{
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = item.Validate()
	if err != nil{
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	orderNumber := item.BeforeSaveItem(server.DB)
	item.OrderNumber = orderNumber
	itemCreated , err := item.SaveItem(server.DB)
	if err != nil {
		formattedError := utils.ErrorType(err.Error())
		utils.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, itemCreated.ID))
	utils.Respond(w, http.StatusOK, itemCreated)
}

func (server *Server) GetItems (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bid, err := strconv.ParseInt(vars["id"], 10,64)
	if err != nil{
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	item := models.Item{}
	items, err :=item.FindAllItem(server.DB, uint32(bid))
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	utils.Respond(w, http.StatusOK, items)
}

func (server *Server) GetItem (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	iid, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	item := models.Item{}
	selectedItem, err := item.FindItem(server.DB, uint32(iid))
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	utils.Respond(w, http.StatusOK, selectedItem)
}

func (server *Server) UpdateItem (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	iid, err := strconv.ParseInt(vars["id"], 10,64)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	item := models.Item{}
	err = server.DB.Model(models.Item{}).Where("id = ?", iid).First(&item).Error
	if err != nil {
		utils.ERROR(w, http.StatusNotFound, errors.New("item not found"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	itemUpdate := models.Item{}
	err = json.Unmarshal(body, &itemUpdate)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = itemUpdate.Validate()
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedItem, err := itemUpdate.UpdateItem(server.DB, uint32(iid))
	if err != nil {
		formattedError := utils.ErrorType(err.Error())
		utils.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	utils.Respond(w, http.StatusOK, updatedItem)
}

func (server *Server) DeleteItem (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	iid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	item := models.Item{}
	err = server.DB.Debug().Model(models.Item{}).Where("id = ?", iid).First(&item).Error
	if err != nil {
		utils.ERROR(w, http.StatusNotFound, errors.New("unauthorized"))
		return
	}
	_, err = item.DeleteItem(server.DB, uint32(iid))
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", iid))
	utils.Respond(w, http.StatusNoContent, "")
}


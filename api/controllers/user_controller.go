package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nozgurozturk/mercury-server/api/models"
	"github.com/nozgurozturk/mercury-server/api/utils"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (server *Server) SignUp(w http.ResponseWriter, r *http.Request){

	user := &models.User{}
	_ = user.BeforeSave()

	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		utils.Respond(w, http.StatusUnprocessableEntity, utils.Message(false, "Invalid request"))
		return
	}

	err = user.Validate()
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	resp, err := user.SaveUser(server.DB) //Create account
	if err != nil{
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	utils.Respond(w, http.StatusOK, resp)
}

func (server *Server) GetUser (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	uid ,err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil{
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user := models.User{}
	selectedUser, err :=user.FindUser(server.DB, uint32(uid))
	if err != nil{
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	utils.Respond(w, http.StatusOK, selectedUser)
}

func (server *Server) UpdateUser (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	uid ,err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil{
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = user.Validate()
	if err != nil{
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedUser, err := user.UpdateUser(server.DB, uint32(uid))
	if err != nil{
		formattedError := utils.ErrorType(err.Error())
		utils.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	utils.Respond(w, http.StatusOK, updatedUser)
}

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	user := models.User{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, err = user.DeleteUser(server.DB, uint32(uid))
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	utils.Respond(w, http.StatusNoContent, "")
}
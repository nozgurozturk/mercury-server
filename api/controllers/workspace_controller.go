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


func (server *Server) CreateWorkspace (w http.ResponseWriter, r *http.Request){


	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		utils.ERROR(w , http.StatusUnprocessableEntity, err)
		return
	}
	user := r.Context().Value("user").(uint32)
	workspace := &models.Workspace{}

	err = json.Unmarshal(body, workspace)
	if err != nil{
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = workspace.Validate()
	if err != nil{
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	workspaceCreated , err := workspace.SaveWorkspace(server.DB, user)
	if err != nil {
		formattedError := utils.ErrorType(err.Error())
		utils.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, workspaceCreated.ID))
	utils.Respond(w, http.StatusOK, workspaceCreated)
}

func (server *Server) GetWorkspaces (w http.ResponseWriter, r *http.Request){
	workspace := &models.Workspace{}
	uid := r.Context().Value("user").(uint32)
	fmt.Println(r)
	boards, err := workspace.FindAllWorkspace(server.DB, uid)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	utils.Respond(w, http.StatusOK, boards)
}

func (server *Server) GetWorkspace (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	//uid := r.Context().Value("user").(uint32)
	wid, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	workspace := models.Workspace{}
	selectedBoard, err := workspace.FindWorkspace(server.DB, uint32(wid))
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	utils.Respond(w, http.StatusOK, selectedBoard)
}

func (server *Server) UpdateWorkspace (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	//uid := r.Context().Value("user").(uint32)

	wid, err := strconv.ParseInt(vars["id"], 10,64)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	workspace := models.Workspace{}
	err = server.DB.Model(models.Workspace{}).Where("id = ?", wid).First(&workspace).Error
	if err != nil {
		utils.ERROR(w, http.StatusNotFound, errors.New("workspace not found"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	workspaceUpdate := models.Workspace{}
	err = json.Unmarshal(body, &workspaceUpdate)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = workspaceUpdate.Validate()
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedWorkspace, err := workspaceUpdate.UpdateWorkspace(server.DB, uint32(wid))
	if err != nil {
		formattedError := utils.ErrorType(err.Error())
		utils.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	utils.Respond(w, http.StatusOK, updatedWorkspace)
}

func (server *Server) DeleteWorkspace (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	//uid := r.Context().Value("user").(uint32)

	wid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	workspace := models.Workspace{}
	err = server.DB.Debug().Model(models.Workspace{}).Where("id = ?", wid).First(&workspace).Error
	if err != nil {
		utils.ERROR(w, http.StatusNotFound, errors.New("unauthorized"))
		return
	}
	_, err = workspace.DeleteWorkspace(server.DB, uint32(wid))
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", wid))
	utils.Respond(w, http.StatusNoContent, "")
}


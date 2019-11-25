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


func (server *Server) CreateBoard (w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		utils.ERROR(w , http.StatusUnprocessableEntity, err)
		return
	}
	//user := r.Context().Value("user").(uint32)
	board := &models.Board{}
	wid, err := strconv.ParseInt(vars["id"], 10, 64)
	board.WorkspaceID = uint32(wid)
	err = json.Unmarshal(body, board)
	if err != nil{
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = board.Validate()
	if err != nil{
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	boardCreated , err := board.SaveBoard(server.DB)
	if err != nil {
		formattedError := utils.ErrorType(err.Error())
		utils.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, boardCreated.ID))
	utils.Respond(w, http.StatusOK, boardCreated)
}

func (server *Server) GetBoards (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	wid, err := strconv.ParseInt(vars["id"], 10, 64)
	board := &models.Board{}
	//uid := r.Context().Value("user").(uint32)
	fmt.Println(r)
	boards, err := board.FindAllBoard(server.DB, uint32(wid))
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	utils.Respond(w, http.StatusOK, boards)
}

func (server *Server) GetBoard (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	uid := r.Context().Value("user").(uint32)
	bid, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	board := models.Board{}
	selectedBoard, err := board.FindBoard(server.DB, uint32(bid), uid)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	utils.Respond(w, http.StatusOK, selectedBoard)
}

func (server *Server) UpdateBoard (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	uid := r.Context().Value("user").(uint32)

	bid, err := strconv.ParseInt(vars["id"], 10,64)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	board := models.Board{}
	err = server.DB.Model(models.Board{}).Where("id = ?", bid).First(&board).Error
	if err != nil {
		utils.ERROR(w, http.StatusNotFound, errors.New("board not found"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	boardUpdate := models.Board{}
	err = json.Unmarshal(body, &boardUpdate)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = boardUpdate.Validate()
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedBoard, err := boardUpdate.UpdateBoard(server.DB, uint32(bid), uid)
	if err != nil {
		formattedError := utils.ErrorType(err.Error())
		utils.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	utils.Respond(w, http.StatusOK, updatedBoard)
}

func (server *Server) DeleteBoard (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	uid := r.Context().Value("user").(uint32)

	bid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	board := models.Board{}
	err = server.DB.Debug().Model(models.Board{}).Where("id = ?", bid).First(&board).Error
	if err != nil {
		utils.ERROR(w, http.StatusNotFound, errors.New("unauthorized"))
		return
	}

	_, err = board.DeleteBoard(server.DB, uint32(bid), uid)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", bid))
	utils.Respond(w, http.StatusNoContent, "")
}


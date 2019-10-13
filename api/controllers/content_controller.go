package controllers

import(
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nozgurozturk/startpage_server/api/auth"
	"github.com/nozgurozturk/startpage_server/api/models"
	"github.com/nozgurozturk/startpage_server/api/utils"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (server *Server) CreateContent (w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		utils.ERROR(w , http.StatusUnprocessableEntity, err)
		return
	}
	content := models.Content{}
	err = json.Unmarshal(body, &content)
	if err != nil{
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = content.Validate()
	if err != nil{
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil{
		utils.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}
	if uid != content.OwnerId{
		utils.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	contentCreated , err := content.SaveContent(server.DB)
	if err != nil {
		formattedError := utils.ErrorType(err.Error())
		utils.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, contentCreated.ID))
	utils.JSON(w, http.StatusOK, contentCreated)
}

func (server *Server) GetContents (w http.ResponseWriter, r *http.Request){
	content := models.Content{}
	contents, err :=content.FindAllContent(server.DB)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	utils.JSON(w, http.StatusOK, contents)
}

func (server *Server) GetContent (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	cid, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	content := models.Content{}
	selectedContent, err := content.FindContent(server.DB, uint32(cid))
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	utils.JSON(w, http.StatusOK, selectedContent)
}

func (server *Server) UpdateContent (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	cid, err := strconv.ParseInt(vars["id"], 10,64)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		utils.ERROR(w , http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}
	content := models.Content{}
	err = server.DB.Model(models.Content{}).Where("id = ?", cid).First(&content).Error
	if err != nil {
		utils.ERROR(w, http.StatusNotFound, errors.New("content not found"))
		return
	}
	if uid != content.OwnerId {
		utils.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	contentUpdate := models.Content{}
	err = json.Unmarshal(body, &contentUpdate)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	if uid != contentUpdate.OwnerId {
		utils.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	err = contentUpdate.Validate()
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedContent, err := contentUpdate.UpdateContent(server.DB, uint32(cid))
	if err != nil {
		formattedError := utils.ErrorType(err.Error())
		utils.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	utils.JSON(w, http.StatusOK, updatedContent)
}

func (server *Server) DeleteContent (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	cid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		utils.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	content := models.Content{}
	err = server.DB.Debug().Model(models.Content{}).Where("id = ?", cid).First(&content).Error
	if err != nil {
		utils.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	if uid != content.OwnerId{
		utils.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	_, err = content.DeleteContent(server.DB, uint32(cid), uid)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", cid))
	utils.JSON(w, http.StatusNoContent, "")
}


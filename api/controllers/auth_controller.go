package controllers

import (
	"encoding/json"
	"github.com/nozgurozturk/startpage_server/api/auth"
	"github.com/nozgurozturk/startpage_server/api/models"
	"github.com/nozgurozturk/startpage_server/api/utils"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
)


func (server *Server) Login(w http.ResponseWriter, r *http.Request){
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
	err = user.Validate("login")
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := utils.ErrorType(err.Error())
		utils.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	utils.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (string, error) {
	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPass(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}
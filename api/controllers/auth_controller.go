package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/nozgurozturk/mercury-server/api/models"
	"github.com/nozgurozturk/mercury-server/api/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request){
	fmt.Println("Welcome to Mercury")
}

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		utils.Respond(w, http.StatusUnprocessableEntity, utils.Message(false, "Invalid request"))
		return
	}
	resp,err := server.SignIn(user.Email, user.Password)

	utils.Respond(w, http.StatusOK, resp)
}

func (server *Server) SignIn(email, password string) (map[string]interface{}, error) {

	user := &models.User{}
	err := server.DB.Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "Email address not found"), err

		}
		return utils.Message(false, "Connection Error"), err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil  { //Password does not match!
		fmt.Println(err)
		return utils.Message(false, "Invalid login credentials"), err
	}
	//Worked! Logged In
	user.Password = ""

	//Create JWT token
	tk := &models.Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("API_SECRET")))
	user.Token = tokenString //Store the token in the response

	resp := utils.Message(true, "Logged In")
	resp["user"] = user
	return resp, err
}




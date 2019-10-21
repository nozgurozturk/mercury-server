package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/nozgurozturk/startpage_server/api/auth"
	"github.com/nozgurozturk/startpage_server/api/models"
	"log"
	"net/http"
	"os"
)
type Server struct{
	DB *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize (){
	var err error
	e := godotenv.Load()
	if e != nil{
		fmt.Print(e)
	}
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	username := os.Getenv("USER_NAME")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbUri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host , port, username, password, dbName)

	server.DB, err = gorm.Open("postgres", dbUri)

	if err != nil{
		fmt.Print(err)
	}
	server.DB.Debug().AutoMigrate(&models.User{}, &models.Board{}, &models.Item{}, &models.Link{})

}

func (server *Server) Run(port string) {
	server.Router = mux.NewRouter()
	server.initializeRoutes()

	server.Router.Use(auth.JwtAuthentication)

	fmt.Println("🚀 on" + port)
	log.Fatal(http.ListenAndServe(port, server.Router))
}



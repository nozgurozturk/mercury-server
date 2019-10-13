package api

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/nozgurozturk/startpage_server/api/controllers"
	"os"
)

var server = controllers.Server{}

func Start() {

	e := godotenv.Load()
	if e != nil{
		fmt.Print(e)
	}
	server.Initialize()

	port := os.Getenv("SERVER_PORT")
	if port == ""{
		port = ":8000"
	}
	server.Run(port)

}

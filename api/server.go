package api

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/nozgurozturk/mercury-server/api/controllers"
	"os"
)

var server = controllers.Server{}

func Start() {

	e := godotenv.Load()
	if e != nil{
		fmt.Print(e)
	}
	server.Initialize()

	port := ":" + os.Getenv("PORT")
	if port == ""{
		port = ":8000"
	}
	server.Run(port)
}

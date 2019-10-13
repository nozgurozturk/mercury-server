package controllers

import (
	"github.com/nozgurozturk/startpage_server/api/middlewares"
)

func (server *Server) initializeRoutes() {
	server.loginRoute()
	server.userRoutes()
	server.contentRoutes()
}

func (server *Server) loginRoute(){

	server.Router.HandleFunc("/login", middlewares.JsonMiddleware(server.Login)).Methods("POST")
}

func (server *Server) userRoutes(){
	server.Router.HandleFunc("/users", middlewares.JsonMiddleware(server.CreateUser)).Methods("POST")
	server.Router.HandleFunc("/users", middlewares.JsonMiddleware(server.GetUsers)).Methods("GET")
	server.Router.HandleFunc("/users/{id}", middlewares.JsonMiddleware(server.GetUser)).Methods("GET")
	server.Router.HandleFunc("/users/{id}", middlewares.JsonMiddleware(middlewares.AuthMiddleware(server.UpdateUser))).Methods("PUT")
	server.Router.HandleFunc("/users/{id}", middlewares.AuthMiddleware(server.DeleteUser)).Methods("DELETE")
}

func (server *Server) contentRoutes(){
	server.Router.HandleFunc("/contents", middlewares.JsonMiddleware(server.CreateContent)).Methods("POST")
	server.Router.HandleFunc("/contents", middlewares.JsonMiddleware(server.GetContents)).Methods("GET")
	server.Router.HandleFunc("/contents/{id}", middlewares.JsonMiddleware(server.GetContent)).Methods("GET")
	server.Router.HandleFunc("/contents/{id}", middlewares.JsonMiddleware(middlewares.AuthMiddleware(server.UpdateContent))).Methods("PUT")
	server.Router.HandleFunc("/contents/{id}", middlewares.AuthMiddleware(server.DeleteContent)).Methods("DELETE")
}
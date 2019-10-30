package controllers

import "net/http"

func (server *Server) initializeRoutes() {
	server.loginRoute()
	server.userRoutes()
	server.boardRoutes()
	server.itemRoutes()
	server.linkRoutes()
}

func (server *Server) loginRoute() {
	server.Router.HandleFunc("/", server.Home).Methods(http.MethodGet)
	server.Router.HandleFunc("/login", server.Login).Methods(http.MethodPost)
	server.Router.HandleFunc("/signup", server.SignUp).Methods(http.MethodPost)

}

func (server *Server) userRoutes() {

	server.Router.HandleFunc("/user/{id}", server.GetUser).Methods(http.MethodGet)
	server.Router.HandleFunc("/user/{id}", server.UpdateUser).Methods(http.MethodPut)
	server.Router.HandleFunc("/user/{id}", server.DeleteUser).Methods(http.MethodDelete)

}

func (server *Server) boardRoutes() {

	server.Router.HandleFunc("/board", server.CreateBoard).Methods(http.MethodPost)
	server.Router.HandleFunc("/board", server.GetBoards).Methods(http.MethodGet)
	server.Router.HandleFunc("/board/{id}", server.GetBoard).Methods(http.MethodGet)
	server.Router.HandleFunc("/board/{id}", server.UpdateBoard).Methods(http.MethodPut)
	server.Router.HandleFunc("/board/{id}", server.DeleteBoard).Methods(http.MethodDelete)

}

func (server *Server) itemRoutes() {

	server.Router.HandleFunc("/item", server.CreateItem).Methods(http.MethodPost)
	server.Router.HandleFunc("/board/{id}/item", server.GetItems).Methods(http.MethodGet)
	server.Router.HandleFunc("/item/{id}", server.GetItem).Methods(http.MethodGet)
	server.Router.HandleFunc("/item/{id}", server.UpdateItem).Methods(http.MethodPut)
	server.Router.HandleFunc("/item/{id}", server.DeleteItem).Methods(http.MethodDelete)

}

func (server *Server) linkRoutes() {

	server.Router.HandleFunc("/link", server.CreateLink).Methods(http.MethodPost)
	server.Router.HandleFunc("/items/{id}/link", server.GetLinks).Methods(http.MethodGet)
	server.Router.HandleFunc("/link/{id}", server.GetLink).Methods(http.MethodGet)
	server.Router.HandleFunc("/link/{id}", server.UpdateLink).Methods(http.MethodPut)
	server.Router.HandleFunc("/link/{id}", server.DeleteLink).Methods(http.MethodDelete)

}

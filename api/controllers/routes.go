package controllers

func (server *Server) initializeRoutes() {
	server.loginRoute()
	server.userRoutes()
	server.boardRoutes()
	server.itemRoutes()
	server.linkRoutes()
}

func (server *Server) loginRoute() {

	server.Router.HandleFunc("/login", server.Login).Methods("POST")
	server.Router.HandleFunc("/signup", server.SignUp).Methods("POST")

}

func (server *Server) userRoutes() {

	server.Router.HandleFunc("/user/{id}", server.GetUser).Methods("GET")
	server.Router.HandleFunc("/user/{id}", server.UpdateUser).Methods("PUT")
	server.Router.HandleFunc("/user/{id}", server.DeleteUser).Methods("DELETE")

}

func (server *Server) boardRoutes() {

	server.Router.HandleFunc("/board", server.CreateBoard).Methods("POST")
	server.Router.HandleFunc("/me/board", server.GetBoards).Methods("GET")
	server.Router.HandleFunc("/me/board/{id}", server.GetBoard).Methods("GET")
	server.Router.HandleFunc("/board/{id}", server.UpdateBoard).Methods("PUT")
	server.Router.HandleFunc("/board/{id}", server.DeleteBoard).Methods("DELETE")

}

func (server *Server) itemRoutes() {

	server.Router.HandleFunc("/item", server.CreateItem).Methods("POST")
	server.Router.HandleFunc("/me/item", server.GetItems).Methods("GET")
	server.Router.HandleFunc("/me/item/{id}", server.GetItem).Methods("GET")
	server.Router.HandleFunc("/item/{id}", server.UpdateItem).Methods("PUT")
	server.Router.HandleFunc("/item/{id}", server.DeleteItem).Methods("DELETE")

}

func (server *Server) linkRoutes() {

	server.Router.HandleFunc("/link", server.CreateLink).Methods("POST")
	server.Router.HandleFunc("/me/link", server.GetLinks).Methods("GET")
	server.Router.HandleFunc("/me/link/{id}", server.GetLink).Methods("GET")
	server.Router.HandleFunc("/link/{id}", server.UpdateLink).Methods("PUT")
	server.Router.HandleFunc("/link/{id}", server.DeleteLink).Methods("DELETE")

}

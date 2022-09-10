package controllers

import "github.com/gin-gonic/gin"

func (server *Server) initializeRoutes() {
	// server.Router = mux.NewRouter()
	// server.Router.HandleFunc("/", server.Home).Methods("GET")
	// server.Router.HandleFunc("/products", server.Products).Methods("GET")

	// staticFileDirectory := http.Dir("./assets/")
	// staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDirectory))
	// server.Router.PathPrefix("/public/").Handler(staticFileHandler).Methods("GET")
	server.Router = gin.Default()
	server.Router.GET("/users", server.Users)
}

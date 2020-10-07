package config

import (
	"pink/app/handlers"
	"pink/app/middlewares"

	"github.com/gorilla/mux"
)

// Router function
func initRouters() *mux.Router {
	router := mux.NewRouter()
	// Init all router in application
	router.HandleFunc("/homes", handlers.HomeHandler)
	router.HandleFunc("/sessions/auth", handlers.Auth).Methods("POST")
	middlewares := middlewares.AuthMiddleware{}
	router.Use(middlewares.AuthAPIToken)
	return router
}

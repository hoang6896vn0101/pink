package config

import (
	"pink/app/handlers"

	"github.com/gorilla/mux"
)

// Router function
func initRouters() *mux.Router {
	router := mux.NewRouter()
	// Init all router in application
	router.HandleFunc("/homes", handlers.HomeHandler)
	return router
}

package config

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// RunBoot func
// Example:
// 1. ENV variables.
// 2. Settings variables.
// 3. Init routes
func runBoot(env string, port string) {
	// Run application on port
	srv := &http.Server{
		Handler:      initRouters(),
		Addr:         fmt.Sprintf("127.0.0.1:%s", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

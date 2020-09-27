package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// HomeHandler func
// Arguments:
// 1. w -> http.ResponseWriter
// 2. r -> *http.Request
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

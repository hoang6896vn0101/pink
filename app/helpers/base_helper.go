package helpers

import (
	"encoding/json"
	"net/http"
)

// FormatJSON func
// Arguments:
// 1. writer -> http.ResponseWriter
// 2. data -> interface
// Return:
// void
func FormatJSON(writer http.ResponseWriter, data interface{}) {
	json, err := json.Marshal(data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(json)
}

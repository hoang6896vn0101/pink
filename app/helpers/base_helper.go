package helpers

import (
	"encoding/json"
	"net/http"
	"pink/app/errors"
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
		handlerError := errors.HandlersError{
			Writer:  writer,
			Message: err.Error(),
			Status:  http.StatusInternalServerError}
		handlerError.InternalServerError()
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(json)
}

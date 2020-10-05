package errors

import (
	"net/http"
	"pink/infrastructure/third_parties/slack"
)

// HandlersError func
type HandlersError struct {
	Writer  http.ResponseWriter
	Message string
	Status  int
}

// InternalServerError func
// Arguments:
// 1. writer -> http.ResponseWriter
// 2. message -> string
func (h *HandlersError) Error() {
	defer slack.PushNotification(h.Message)
	http.Error(h.Writer, h.Message, h.Status)
	panic(h.Message)
}

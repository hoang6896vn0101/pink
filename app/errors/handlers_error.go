package errors

import (
	"net/http"
	protocols "pink/protocols/third_parties"
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
func (h *HandlersError) InternalServerError() {
	slack := protocols.Slack{}
	defer slack.PushNotification(h.Message)
	http.Error(h.Writer, h.Message, h.Status)
	panic(h.Message)
}

// AuthError func
// Arguments:
// 1. writer -> http.ResponseWriter
// 2. message -> string
func (h *HandlersError) AuthError() {
	http.Error(h.Writer, h.Message, h.Status)
}

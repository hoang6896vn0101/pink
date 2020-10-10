package handlers

import (
	"fmt"
	"net/http"
	"pink/domains/services"
)

// Auth func
func Auth(writer http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	service := services.UserManagerService{}
	if service.Login(email) {
		// Handler response return to cliend
		fmt.Print("Login Success")
	} else {
		fmt.Print("Login Fail")
		// Handler response return to cliend
	}
}

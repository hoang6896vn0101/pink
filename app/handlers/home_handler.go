package handlers

import (
	"net/http"
	"pink/app/helpers"
	"pink/domains/entities"
	"pink/domains/models"
)

// HomeHandler func
// Arguments:
// 1. w -> http.ResponseWriter
// 2. r -> *http.Request
func HomeHandler(writer http.ResponseWriter, response *http.Request) {
	users := models.UserModel{
		Account: entities.AccountEntity{ID: 1, UserName: "hoang6896vn0101", PassWord: "Default"},
		User:    entities.UserEntity{ID: 1, Name: "Hoang Pham Dang"},
	}
	helpers.FormatJSON(writer, users)
}

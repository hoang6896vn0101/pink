package model_repos

import (
	e "pink/domains/entities"
	u "pink/domains/models"
)

type UserModel u.UserModel

type UserModelRepo interface {
	GetAllUser() []UserModel
}

func (u UserModel) GetAllUser() []UserModel {
	// Call to Database and run
	account := e.AccountEntity{ID: 1, UserName: "foo_1996", PassWord: "Aa@123456"}
	user := e.UserEntity{ID: 1, Name: "foo"}
	users := []UserModel{{Account: account, User: user}}
	return users
}

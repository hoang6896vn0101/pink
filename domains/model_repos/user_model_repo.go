package modelrepos

import (
	e "pink/domains/entities"
	u "pink/domains/models"
)

// UserModel model
type UserModel u.UserModel

// UserModelRepoInf interface
type UserModelRepoInf interface {
	GetAllUser() []UserModel
	Authenticate(apiToken string) bool
	MatchEmail(email string) bool
	SendAPIToken(email string) bool
}

// UserModelRepo struct
type UserModelRepo struct {
}

// NewUserRepo func
func NewUserRepo() *UserModelRepo {
	return &UserModelRepo{}
}

// GetAllUser func
func (userRepo UserModelRepo) GetAllUser() []UserModel {
	// Call to Database and run
	account := e.AccountEntity{ID: 1, UserName: "foo_1996", PassWord: "Aa@123456"}
	user := e.UserEntity{ID: 1, Name: "foo"}
	users := []UserModel{{Account: account, User: user}}
	return users
}

// Authenticate func
func (userRepo UserModelRepo) Authenticate(apiToken string) bool {
	account := e.AccountEntity{ID: 1, UserName: "foo_1996", PassWord: "Aa@123456", APIToken: "1234"}
	if apiToken == account.APIToken {
		return true
	}
	return false
}

// MatchEmail func
func (userRepo UserModelRepo) MatchEmail(email string) bool {
	return true
}

// SendAPIToken func
func (userRepo UserModelRepo) SendAPIToken(email string) bool {
	return true
}

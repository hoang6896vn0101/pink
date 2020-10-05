package models

import (
	e "pink/domains/entities"
)

// UserModel struct
type UserModel struct {
	Account e.AccountEntity
	User    e.UserEntity
}

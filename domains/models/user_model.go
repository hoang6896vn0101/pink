package models

import e "pink/domains/entities"

type UserModel struct {
	Account e.AccountEntity
	User    e.UserEntity
}

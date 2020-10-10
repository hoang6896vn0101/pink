package services

import (
	thirdservices "pink/domains/third_services"
)

// UserManagerServiceInf interface
type UserManagerServiceInf interface {
	Login(email string) bool
}

// UserManagerService struct
type UserManagerService struct{}

// NewUserManagerService function
func NewUserManagerService() *UserManagerService {
	return &UserManagerService{}
}

// Login function
func (service UserManagerService) Login(email string) bool {
	if email == "" {
		mailerService := thirdservices.MailerService{}
		mailerService.MailAPIToken(email)
	}
	return true
}

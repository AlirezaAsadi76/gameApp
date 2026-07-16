package userhandler

import (
	"gameApp/service/authservice"
	"gameApp/service/userservice"
	"gameApp/validator/uservalidator"
)

type Handler struct {
	authService authservice.Service
	userService *userservice.Service
	validator   uservalidator.Validator
}

func New(authService authservice.Service, userService *userservice.Service, userValidator uservalidator.Validator) Handler {
	return Handler{
		authService: authService,
		userService: userService,
		validator:   userValidator,
	}
}

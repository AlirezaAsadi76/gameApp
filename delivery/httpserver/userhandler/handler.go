package userhandler

import (
	"gameApp/service/authservice"
	"gameApp/service/userservice"
	"gameApp/validator/uservalidator"
)

type Handler struct {
	authConfig  authservice.Config
	authService authservice.Service
	userService *userservice.Service
	validator   uservalidator.Validator
}

func New(authConfig authservice.Config, authService authservice.Service, userService *userservice.Service, userValidator uservalidator.Validator) Handler {
	return Handler{
		authConfig:  authConfig,
		authService: authService,
		userService: userService,
		validator:   userValidator,
	}
}

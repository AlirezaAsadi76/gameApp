package httpserver

import (
	"fmt"
	"gameApp/config"
	"gameApp/service/authservice"
	"gameApp/service/userservice"
	"gameApp/validator/uservalidator"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Server struct {
	config      config.Config
	authService authservice.Service
	userService *userservice.Service
	validator   uservalidator.Validator
}

func New(config config.Config, authService authservice.Service, userService *userservice.Service, userValidator uservalidator.Validator) Server {
	return Server{
		config:      config,
		authService: authService,
		userService: userService,
		validator:   userValidator,
	}
}

func (s Server) Start() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	userGroup := e.Group("/users")
	userGroup.POST("/register", s.userRegister)
	userGroup.POST("/login", s.userLogin)
	userGroup.GET("/profile", s.userProfile)

	if err := e.Start(fmt.Sprintf(":%d", s.config.HttpServer.Port)); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}

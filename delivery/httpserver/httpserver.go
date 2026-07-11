package httpserver

import (
	"fmt"
	"gameApp/config"
	"gameApp/service/authservice"
	"gameApp/service/userservice"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Server struct {
	config      config.Config
	authService authservice.Service
	userService *userservice.Service
}

func New(config config.Config, authService authservice.Service, userService *userservice.Service) Server {
	return Server{
		config:      config,
		authService: authService,
		userService: userService,
	}
}

func (s Server) Start() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	e.POST("/users/register", s.userRegister)

	if err := e.Start(fmt.Sprintf(":%d", s.config.HttpServer.Port)); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}

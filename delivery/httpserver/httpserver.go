package httpserver

import (
	"fmt"
	"gameApp/config"
	"gameApp/delivery/httpserver/userhandler"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Server struct {
	config      config.Config
	userHandler userhandler.Handler
}

func New(config config.Config, userHandler userhandler.Handler) Server {
	return Server{
		config:      config,
		userHandler: userHandler,
	}
}

func (s Server) Start() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	s.userHandler.SetRoutes(e)

	if err := e.Start(fmt.Sprintf(":%d", s.config.HttpServer.Port)); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}

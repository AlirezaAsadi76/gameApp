package userhandler

import (
	"gameApp/delivery/httpserver/middleware"

	"github.com/labstack/echo/v5"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")

	userGroup.POST("/register", h.userRegister, middleware.Auth(h.authConfig, h.authService))
	userGroup.POST("/login", h.userLogin)
	userGroup.GET("/profile", h.userProfile)
}

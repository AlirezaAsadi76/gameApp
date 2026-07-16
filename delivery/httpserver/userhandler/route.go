package userhandler

import "github.com/labstack/echo/v5"

func (h Handler) SetRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")
	userGroup.POST("/register", h.userRegister)
	userGroup.POST("/login", h.userLogin)
	userGroup.GET("/profile", h.userProfile)
}

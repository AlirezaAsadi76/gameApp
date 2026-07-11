package httpserver

import (
	"gameApp/service/userservice"
	"net/http"

	"github.com/labstack/echo/v5"
)

func (s Server) userRegister(c *echo.Context) error {

	var registerRequest userservice.RegisterRequest
	bErr := c.Bind(&registerRequest)
	if bErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, bErr.Error())
	}

	user, rErr := s.userService.Register(registerRequest)
	if rErr != nil {

		return echo.NewHTTPError(http.StatusBadRequest, rErr.Error())
	}
	return c.JSON(http.StatusCreated, user)
}

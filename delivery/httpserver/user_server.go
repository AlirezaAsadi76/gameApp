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

func (s Server) userLogin(c *echo.Context) error {

	var loginRequest userservice.LoginRequest
	if err := c.Bind(&loginRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userResponse, err := s.userService.Login(loginRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, userResponse)
}

func (s Server) userProfile(c *echo.Context) error {

	authorization := c.Request().Header.Get("Authorization")

	claims, pErr := s.authService.ParseToken(authorization, s.config.Auth.SignKey)
	if pErr != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, pErr.Error())
	}

	profileResponse, err := s.userService.Profile(userservice.ProfileRequest{
		UserId: claims.UserID,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, profileResponse)
}

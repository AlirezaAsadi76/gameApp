package httpserver

import (
	"gameApp/pkg/httpmsg"
	"gameApp/service/userservice"
	"net/http"

	"github.com/labstack/echo/v5"
)

func (s Server) userRegister(c *echo.Context) error {

	var registerRequest userservice.RegisterRequest
	bErr := c.Bind(&registerRequest)
	if bErr != nil {
		msg, code := httpmsg.CodeAndMessage(bErr)
		return echo.NewHTTPError(code, msg)
	}

	user, rErr := s.userService.Register(registerRequest)
	if rErr != nil {
		msg, code := httpmsg.CodeAndMessage(rErr)
		return echo.NewHTTPError(code, msg)
	}
	return c.JSON(http.StatusCreated, user)
}

func (s Server) userLogin(c *echo.Context) error {

	var loginRequest userservice.LoginRequest
	if err := c.Bind(&loginRequest); err != nil {
		msg, code := httpmsg.CodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}

	userResponse, err := s.userService.Login(loginRequest)
	msg, code := httpmsg.CodeAndMessage(err)
	if err != nil {
		return echo.NewHTTPError(code, msg)
	}
	return c.JSON(http.StatusOK, userResponse)
}

func (s Server) userProfile(c *echo.Context) error {

	authorization := c.Request().Header.Get("Authorization")

	claims, pErr := s.authService.ParseToken(authorization, s.config.Auth.SignKey)
	msg, code := httpmsg.CodeAndMessage(pErr)
	if pErr != nil {
		return echo.NewHTTPError(code, msg)
	}

	profileResponse, err := s.userService.Profile(userservice.ProfileRequest{
		UserId: claims.UserID,
	})
	msg, code = httpmsg.CodeAndMessage(err)
	if err != nil {
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, profileResponse)
}

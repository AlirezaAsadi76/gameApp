package userhandler

import (
	"gameApp/params"
	"gameApp/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v5"
)

func (h Handler) userProfile(c *echo.Context) error {

	authorization := c.Request().Header.Get("Authorization")

	claims, pErr := h.authService.ParseToken(authorization)
	msg, code := httpmsg.CodeAndMessage(pErr)
	if pErr != nil {
		return echo.NewHTTPError(code, msg)
	}

	profileResponse, err := h.userService.Profile(params.ProfileRequest{
		UserId: claims.UserID,
	})
	msg, code = httpmsg.CodeAndMessage(err)
	if err != nil {
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, profileResponse)
}

package userhandler

import (
	"gameApp/params"
	"gameApp/pkg/constant"
	"gameApp/pkg/httpmsg"
	"gameApp/service/authservice"
	"net/http"

	"github.com/labstack/echo/v5"
)

func getClaims(c *echo.Context) *authservice.Claims {
	return c.Get(constant.AuthMiddlewareContextKey).(*authservice.Claims)
}

func (h Handler) userProfile(c *echo.Context) error {

	claims := getClaims(c)
	profileResponse, err := h.userService.Profile(params.ProfileRequest{
		UserId: claims.UserID,
	})
	msg, code := httpmsg.CodeAndMessage(err)
	if err != nil {
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, profileResponse)
}

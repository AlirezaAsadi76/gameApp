package userhandler

import (
	"gameApp/params"
	"gameApp/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v5"
)

func (h Handler) userLogin(c *echo.Context) error {

	var loginRequest params.LoginRequest
	if err := c.Bind(&loginRequest); err != nil {
		msg, code := httpmsg.CodeAndMessage(err)
		return echo.NewHTTPError(code, msg)
	}

	fieldError, vErr := h.validator.ValidatorLoginRequest(loginRequest)

	if vErr != nil {

		msg, code := httpmsg.CodeAndMessage(vErr)

		return c.JSON(code, map[string]interface{}{

			"msg":         msg,
			"fieldErrors": fieldError,
		})

	}

	userResponse, err := h.userService.Login(loginRequest)
	msg, code := httpmsg.CodeAndMessage(err)
	if err != nil {
		return echo.NewHTTPError(code, msg)
	}
	return c.JSON(http.StatusOK, userResponse)
}

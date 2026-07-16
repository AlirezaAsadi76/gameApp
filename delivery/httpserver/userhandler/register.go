package userhandler

import (
	"gameApp/params"
	"gameApp/pkg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v5"
)

func (h Handler) userRegister(c *echo.Context) error {

	var registerRequest params.RegisterRequest
	bErr := c.Bind(&registerRequest)

	fieldError, vErr := h.validator.ValidatorRegisterRequest(registerRequest)

	if vErr != nil {

		msg, code := httpmsg.CodeAndMessage(vErr)

		return c.JSON(code, map[string]interface{}{

			"msg":         msg,
			"fieldErrors": fieldError,
		})

	}

	if bErr != nil {
		msg, code := httpmsg.CodeAndMessage(bErr)
		return echo.NewHTTPError(code, msg)
	}

	user, rErr := h.userService.Register(registerRequest)
	if rErr != nil {
		msg, code := httpmsg.CodeAndMessage(rErr)
		return echo.NewHTTPError(code, msg)
	}
	return c.JSON(http.StatusCreated, user)
}

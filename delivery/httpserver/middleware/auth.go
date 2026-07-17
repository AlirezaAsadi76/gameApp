package middleware

import (
	"gameApp/pkg/constant"
	"gameApp/service/authservice"

	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
)

func Auth(authConfig authservice.Config, authService authservice.Service) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		ContextKey:    constant.AuthMiddlewareContextKey,
		SigningKey:    authConfig.SignKey,
		SigningMethod: "HS256",
		ParseTokenFunc: func(c *echo.Context, auth string) (interface{}, error) {
			claims, pErr := authService.ParseToken(auth)
			if pErr != nil {
				return nil, pErr
			}
			return claims, nil
		},
	})
}

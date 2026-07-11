package authservice

import (
	"errors"
	"gameApp/entity"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Config struct {
	SignKey              string
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
	AccessSubject        string
	RefreshSubject       string
}

type Service struct {
	config Config
}

func New(cfg Config) Service {
	return Service{
		config: cfg,
	}
}

func (s Service) CreateAccessToken(user entity.User) (string, error) {
	return createToken(user.ID, s.config.SignKey, s.config.AccessSubject, s.config.AccessTokenDuration)
}

func (s Service) CreateRefreshToken(user entity.User) (string, error) {
	return createToken(user.ID, s.config.SignKey, s.config.RefreshSubject, s.config.RefreshTokenDuration)
}

func createToken(userID uint, signKey, subject string, duration time.Duration) (string, error) {

	t := jwt.New(jwt.GetSigningMethod("HS256"))

	t.Claims = &Claims{
		jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
		userID,
	}

	return t.SignedString([]byte(signKey))
}

func (s Service) ParseToken(tokenStr string, signKey string) (Claims, error) {

	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil
	})

	if err != nil {
		return Claims{}, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return Claims{}, errors.New("invalid token")
	}
	return *claims, nil
}

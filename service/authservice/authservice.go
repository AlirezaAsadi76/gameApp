package authservice

import (
	"errors"
	"gameApp/entity"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
	signKey              string
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
	accessSubject        string
	refreshSubject       string
}

func New(signKey, accessSubject, refreshSubject string, accessTokenDuration, refreshTokenDuration time.Duration) Service {
	return Service{
		signKey:              signKey,
		accessTokenDuration:  accessTokenDuration,
		refreshTokenDuration: refreshTokenDuration,
		accessSubject:        accessSubject,
		refreshSubject:       refreshSubject,
	}
}

func (s Service) CreateAccessToken(user entity.User) (string, error) {
	return createToken(user.ID, s.signKey, s.accessSubject, s.accessTokenDuration)
}

func (s Service) CreateRefreshToken(user entity.User) (string, error) {
	return createToken(user.ID, s.signKey, s.refreshSubject, s.refreshTokenDuration)
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

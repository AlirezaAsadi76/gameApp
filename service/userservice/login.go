package userservice

import (
	"errors"
	"fmt"
	"gameApp/params"
	"gameApp/pkg/hashing"
)

func (s *Service) Login(req params.LoginRequest) (params.LoginResponse, error) {
	// check existences of phone_number from repository
	// get the user by phoneNumber
	//TODO - better use two method
	user, gErr := s.repository.GetUserByPhoneNumber(req.PhoneNumber)
	if gErr != nil {
		return params.LoginResponse{}, fmt.Errorf("unexpected error : %w", gErr)
	}

	if !hashing.CheckPasswordHash(req.Password, user.Password) {
		return params.LoginResponse{}, errors.New("username or password isn't exist")
	}

	accessToken, tErr := s.authGenerator.CreateAccessToken(user)
	if tErr != nil {
		return params.LoginResponse{}, fmt.Errorf("unexpected error : %w", tErr)
	}
	refreshToken, rErr := s.authGenerator.CreateRefreshToken(user)
	if rErr != nil {
		return params.LoginResponse{}, fmt.Errorf("unexpected error : %w", rErr)
	}

	return params.LoginResponse{Tokens: params.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, User: params.UserInfo{
		ID:          user.ID,
		PhoneNumber: user.PhoneNumber,
		Name:        user.Name,
	}}, nil
}

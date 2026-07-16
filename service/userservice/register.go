package userservice

import (
	"fmt"
	"gameApp/entity"
	"gameApp/params"
	"gameApp/pkg/hashing"
)

func (s *Service) Register(req params.RegisterRequest) (params.RegisterResponse, error) {
	// TODO - verify phonenumber

	hash, hErr := hashing.HashPassword(req.Password)
	if hErr != nil {
		return params.RegisterResponse{}, hErr
	}
	// create user in storage
	user := entity.User{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		ID:          0,
		Password:    hash,
	}

	userCreated, rErr := s.repository.Register(user)
	if rErr != nil {
		return params.RegisterResponse{}, fmt.Errorf("unexpected error : %w", rErr)
	}

	// return created user

	return params.RegisterResponse{User: params.UserInfo{
		ID:          userCreated.ID,
		PhoneNumber: userCreated.PhoneNumber,
		Name:        userCreated.Name,
	}}, nil

}

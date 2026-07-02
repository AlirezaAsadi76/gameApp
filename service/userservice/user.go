package userservice

import (
	"errors"
	"fmt"
	"gameApp/entity"
	"gameApp/pkg/phonenumber"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
}

type Service struct {
	repository Repository
}

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type RegisterResponse struct {
	User entity.User
}

func New(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Register(req RegisterRequest) (RegisterResponse, error) {
	// TODO - verify phonenumber

	//validate phone number
	if !phonenumber.IsValid(req.PhoneNumber) {
		return RegisterResponse{}, errors.New("invalid phone")
	}

	// check uniqueness phone number
	if ok, isErr := s.repository.IsPhoneNumberUnique(req.PhoneNumber); isErr != nil || !ok {
		if isErr != nil {
			return RegisterResponse{}, fmt.Errorf("unexpected error : %w", isErr)
		}
		return RegisterResponse{}, fmt.Errorf("phoneNumber is not unique")

	}

	// validate name
	if len(req.Name) < 3 {
		return RegisterResponse{}, errors.New("name is too short, must be 3 characters long")
	}
	// create user in storage
	user := entity.User{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		ID:          0,
	}

	userCreated, rErr := s.repository.Register(user)
	if rErr != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error : %w", rErr)
	}

	// return created user

	return RegisterResponse{userCreated}, nil

}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct{}

func (s *Service) Login(req LoginRequest) (LoginResponse, error) {
	// check existences of phone_number from repository

	// get the user by phoneNumber

	// compare user.password with the req.password

	// return loginResponse
}

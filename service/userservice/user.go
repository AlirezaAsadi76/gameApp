package userservice

import (
	"errors"
	"fmt"
	"gameApp/entity"
	"gameApp/pkg/hashing"
	"gameApp/pkg/phonenumber"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(number string) (entity.User, bool, error)
	GetUserByID(id uint) (entity.User, error)
}

type Service struct {
	repository Repository
}

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
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

	// TODO - check the  password with regex pattern
	// validate password
	if len(req.Password) < 8 {
		return RegisterResponse{}, errors.New("password is too short, must be 8 characters long")
	}

	hash, hErr := hashing.HashPassword(req.Password)
	if hErr != nil {
		return RegisterResponse{}, hErr
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
		return RegisterResponse{}, fmt.Errorf("unexpected error : %w", rErr)
	}

	// return created user

	return RegisterResponse{userCreated}, nil

}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	User entity.User
}

func (s *Service) Login(req LoginRequest) (LoginResponse, error) {
	// check existences of phone_number from repository
	// get the user by phoneNumber
	//TODO - better use two method
	user, exist, gErr := s.repository.GetUserByPhoneNumber(req.PhoneNumber)
	if gErr != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error : %w", gErr)
	}
	if !exist {
		return LoginResponse{}, errors.New("username or password isn't exist")
	}
	// compare user.password with the req.password
	if !hashing.CheckPasswordHash(req.Password, user.Password) {
		return LoginResponse{}, errors.New("username or password isn't exist")
	}

	// return loginResponse

	return LoginResponse{User: user}, nil
}

type ProfileRequest struct {
	UserId uint `json:"user_id"`
}
type ProfileResponse struct {
	Name string `json:"name"`
}

// all request intructor/service should be sanitized

func (s *Service) Profile(req ProfileRequest) (ProfileResponse, error) {

	//TODO - we can use rich error
	user, gErr := s.repository.GetUserByID(req.UserId)
	if gErr != nil {
		return ProfileResponse{}, fmt.Errorf("unexpected error : %w", gErr)
	}
	return ProfileResponse{Name: user.Name}, nil
}

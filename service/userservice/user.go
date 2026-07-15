package userservice

import (
	"errors"
	"fmt"
	"gameApp/dto"
	"gameApp/entity"
	"gameApp/pkg/hashing"
	"gameApp/pkg/richerror"
)

type Repository interface {
	Register(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(number string) (entity.User, bool, error)
	GetUserByID(id uint) (entity.User, error)
}

type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}
type Service struct {
	authGenerator AuthGenerator
	repository    Repository
}

func New(repository Repository, authGenerator AuthGenerator) *Service {
	return &Service{repository: repository, authGenerator: authGenerator}
}

func (s *Service) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {
	// TODO - verify phonenumber

	hash, hErr := hashing.HashPassword(req.Password)
	if hErr != nil {
		return dto.RegisterResponse{}, hErr
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
		return dto.RegisterResponse{}, fmt.Errorf("unexpected error : %w", rErr)
	}

	// return created user

	return dto.RegisterResponse{User: dto.UserInfo{
		ID:          userCreated.ID,
		PhoneNumber: userCreated.PhoneNumber,
		Name:        userCreated.Name,
	}}, nil

}

func (s *Service) Login(req dto.LoginRequest) (dto.LoginResponse, error) {
	// check existences of phone_number from repository
	// get the user by phoneNumber
	//TODO - better use two method
	user, exist, gErr := s.repository.GetUserByPhoneNumber(req.PhoneNumber)
	if gErr != nil {
		return dto.LoginResponse{}, fmt.Errorf("unexpected error : %w", gErr)
	}
	if !exist {
		return dto.LoginResponse{}, errors.New("username or password isn't exist")
	}
	// compare user.password with the req.password
	if !hashing.CheckPasswordHash(req.Password, user.Password) {
		return dto.LoginResponse{}, errors.New("username or password isn't exist")
	}

	accessToken, tErr := s.authGenerator.CreateAccessToken(user)
	if tErr != nil {
		return dto.LoginResponse{}, fmt.Errorf("unexpected error : %w", tErr)
	}
	refreshToken, rErr := s.authGenerator.CreateRefreshToken(user)
	if rErr != nil {
		return dto.LoginResponse{}, fmt.Errorf("unexpected error : %w", rErr)
	}

	return dto.LoginResponse{Tokens: dto.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, User: dto.UserInfo{
		ID:          user.ID,
		PhoneNumber: user.PhoneNumber,
		Name:        user.Name,
	}}, nil
}

// all request intructor/service should be sanitized

func (s *Service) Profile(req dto.ProfileRequest) (dto.ProfileResponse, error) {
	const Op = "UserService.Profile"
	//TODO - we can use rich error
	userSelected, gErr := s.repository.GetUserByID(req.UserId)
	if gErr != nil {
		return dto.ProfileResponse{}, richerror.New(Op).WithError(gErr)
	}

	return dto.ProfileResponse{User: dto.UserInfo{
		ID:          userSelected.ID,
		PhoneNumber: userSelected.PhoneNumber,
		Name:        userSelected.Name,
	}}, nil
}

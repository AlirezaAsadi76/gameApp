package userservice

import (
	"gameApp/entity"
)

type Repository interface {
	Register(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(number string) (entity.User, error)
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

package uservalidator

import "gameApp/entity"

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	GetUserByPhoneNumber(number string) (entity.User, error)
}

const (
	regexPatternPhoneNumber = "^09[0-9]{9}$"
	regexPatternPassword    = `^([a-zA-Z0-9!@#$%^&*]).{8,}$`
)

type Validator struct {
	repository Repository
}

func New(repository Repository) Validator {
	return Validator{
		repository: repository,
	}
}

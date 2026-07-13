package uservalidator

import (
	"errors"
	"fmt"
	"gameApp/dto"
	"gameApp/pkg/msgerror"
	"gameApp/pkg/richerror"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
}
type Validator struct {
	repository Repository
}

func New(repository Repository) Validator {
	return Validator{
		repository: repository,
	}
}

func (v Validator) ValidatorRegisterRequest(req dto.RegisterRequest) (map[string]string, error) {
	const Op = "validator.ValidatorRegisterRequest"

	vErr := validation.ValidateStruct(&req,

		validation.Field(&req.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&req.Password, validation.Required,
			validation.Match(regexp.MustCompile(`^([a-zA-Z0-9!@#$%^&*]).{8,}$`))),
		validation.Field(&req.PhoneNumber, validation.Required,
			validation.Match(regexp.MustCompile("^09[0-9]{9}$")),
			validation.By(v.checkIsPhoneNumberUnique)),
	)
	if vErr != nil {

		fieldErrors := make(map[string]string)
		var errV validation.Errors
		ok := errors.As(vErr, &errV)
		if ok {
			for key, val := range errV {
				fieldErrors[key] = val.Error()
			}
		}
		return fieldErrors, richerror.New(Op).
			WithMessage(msgerror.ErrorMsgInputInValid).
			WithKind(richerror.KindInvalid).
			WithError(vErr).
			WithMeta(map[string]interface{}{"request": req})
	}
	return nil, nil
}

func (v Validator) checkIsPhoneNumberUnique(value interface{}) error {
	phoneNumber := value.(string)

	if ok, isErr := v.repository.IsPhoneNumberUnique(phoneNumber); isErr != nil || !ok {
		if isErr != nil {
			return isErr
		}
		return fmt.Errorf(msgerror.ErrorMsgPhoneNumberIsNotUnique)

	}
	return nil
}

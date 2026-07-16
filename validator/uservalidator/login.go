package uservalidator

import (
	"errors"
	"fmt"
	"gameApp/params"
	"gameApp/pkg/msgerror"
	"gameApp/pkg/richerror"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

func (v Validator) ValidatorLoginRequest(req params.LoginRequest) (map[string]string, error) {
	const Op = "validator.ValidatorLoginRequest"

	vErr := validation.ValidateStruct(&req,
		validation.Field(&req.Password,
			validation.Required,
			validation.Match(regexp.MustCompile(regexPatternPassword))),
		validation.Field(&req.PhoneNumber, validation.Required,
			validation.Match(regexp.MustCompile(regexPatternPhoneNumber)).Error(msgerror.ErrorMsgPhoneNumberIsNotValid),
			validation.By(v.dosePhoneNumberExists)),
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
func (v Validator) dosePhoneNumberExists(value interface{}) error {
	phoneNumber := value.(string)
	_, isErr := v.repository.GetUserByPhoneNumber(phoneNumber)

	if isErr != nil {
		return fmt.Errorf(msgerror.ErrorMsgNotFound)
	}

	return nil
}

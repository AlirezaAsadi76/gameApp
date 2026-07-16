package userservice

import (
	"gameApp/params"
	"gameApp/pkg/richerror"
)

// all request intructor/service should be sanitized

func (s *Service) Profile(req params.ProfileRequest) (params.ProfileResponse, error) {
	const Op = "UserService.Profile"
	//TODO - we can use rich error
	userSelected, gErr := s.repository.GetUserByID(req.UserId)
	if gErr != nil {
		return params.ProfileResponse{}, richerror.New(Op).WithError(gErr)
	}

	return params.ProfileResponse{User: params.UserInfo{
		ID:          userSelected.ID,
		PhoneNumber: userSelected.PhoneNumber,
		Name:        userSelected.Name,
	}}, nil
}

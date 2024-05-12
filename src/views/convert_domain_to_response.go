package view

import (
	"github.com/setxpro/crud-go/src/controllers/model/response"
	model "github.com/setxpro/crud-go/src/models"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}

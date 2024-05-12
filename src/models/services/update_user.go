package service

import (
	resterr "github.com/setxpro/crud-go/src/configurations/rest_err"
	model "github.com/setxpro/crud-go/src/models"
)

func (*userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *resterr.RestErr {
	return nil
}

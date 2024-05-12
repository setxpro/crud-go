package service

import (
	resterr "github.com/setxpro/crud-go/src/configurations/rest_err"
	model "github.com/setxpro/crud-go/src/models"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *resterr.RestErr
	UpdateUser(string, model.UserDomainInterface) *resterr.RestErr
	FindUser(string) (*model.UserDomainInterface, *resterr.RestErr)
	DeleteUser(string) *resterr.RestErr
}

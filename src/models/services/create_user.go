package service

import (
	"fmt"

	"github.com/setxpro/crud-go/src/configurations/logger"
	resterr "github.com/setxpro/crud-go/src/configurations/rest_err"
	model "github.com/setxpro/crud-go/src/models"
	"go.uber.org/zap"
)

// Regra de negócio de toda a aplicação
func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *resterr.RestErr {
	logger.Info("Init CreateUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())
	logger.Info("User created successfully", zap.String("journey", "createUser"))
	return nil
}

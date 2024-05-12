package model

import (
	"fmt"

	"github.com/setxpro/crud-go/src/configurations/logger"
	resterr "github.com/setxpro/crud-go/src/configurations/rest_err"
	"go.uber.org/zap"
)

// Regra de negócio de toda a aplicação
func (ud *UserDomain) CreateUser() *resterr.RestErr {
	logger.Info("Init CreateUser model", zap.String("journey", "createUser"))
	ud.EncryptPassword()

	fmt.Println(ud)
	logger.Info("User created successfully", zap.String("journey", "createUser"))
	return nil
}

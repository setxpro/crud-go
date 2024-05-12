package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/setxpro/crud-go/src/configurations/logger"
	validation "github.com/setxpro/crud-go/src/configurations/validations"
	"github.com/setxpro/crud-go/src/controllers/model/request"
	model "github.com/setxpro/crud-go/src/models"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func InsertUser(c *gin.Context) {

	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))

	// Get object user
	var userRequest request.UserRequest

	// c ->
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error tring to validate user info : ", err, zap.String("journey", "createUser"))

		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}
	logger.Info("User created successfully", zap.String("journey", "createUser"))

	c.String(http.StatusOK, "")
}

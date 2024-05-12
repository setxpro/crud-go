package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/setxpro/crud-go/src/configurations/logger"
	validation "github.com/setxpro/crud-go/src/configurations/validations"
	"github.com/setxpro/crud-go/src/controllers/model/request"
)

func InsertUser(c *gin.Context) {

	logger.Info("Init CreateUser controller")

	// Get object user
	var userRequest request.UserRequest

	// c ->
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error tring to validate user info : ", err)
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}
	logger.Info("User created successfully")
	fmt.Println(userRequest)
}

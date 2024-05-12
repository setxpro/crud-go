package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	validation "github.com/setxpro/crud-go/src/configurations/validations"
	"github.com/setxpro/crud-go/src/controllers/model/request"
)

func InsertUser(c *gin.Context) {
	// Get object user
	var userRequest request.UserRequest

	// c ->
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
}

package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	resterr "github.com/setxpro/crud-go/src/configurations/rest_err"
	"github.com/setxpro/crud-go/src/controllers/model/request"
)

func InsertUser(c *gin.Context) {
	// Get object user
	var userRequest request.UserRequest

	// c ->
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := resterr.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error=%s", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}

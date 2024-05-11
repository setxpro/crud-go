package controller

import (
	"github.com/gin-gonic/gin"
	resterr "github.com/setxpro/crud-go/src/configurations/rest_err"
)

func InsertUser(c *gin.Context) {
	err := resterr.NewBadRequestError("Você chamou a rota de forma inválida")

	c.JSON(err.Code, err)
}

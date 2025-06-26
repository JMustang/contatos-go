package controller

import (
	errorhandler "github.com/JMustang/contatos-go/src/config/error_handler"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := errorhandler.NewBadRequestError("❌ bad request")
	c.JSON(err.Code, err)
}

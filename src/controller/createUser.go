package controller

import (
	"fmt"

	errorhandler "github.com/JMustang/contatos-go/src/config/error_handler"
	"github.com/JMustang/contatos-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		rest_error := errorhandler.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error=%s", err.Error()),
		)

		c.JSON(rest_error.Code, rest_error)
		return
	}

	// TODO: Implementar lógica de criação do usuário
	c.JSON(201, gin.H{"message": "User created successfully"})
}

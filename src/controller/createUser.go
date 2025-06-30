package controller

import (
	"github.com/JMustang/contatos-go/src/config/validation"
	"github.com/JMustang/contatos-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		rest_error := validation.ValidateUserError(err)

		c.JSON(rest_error.Code, rest_error)
		return
	}

	// TODO: Implementar lógica de criação do usuário
	c.JSON(201, gin.H{"message": "User created successfully"})
}

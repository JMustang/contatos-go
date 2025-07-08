package model

import (
	errorhandler "github.com/JMustang/contatos-go/src/config/error_handler"
	"github.com/JMustang/contatos-go/src/config/logger"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *errorhandler.HandlerErr {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	ud.EncryptPassword()

	return nil
}

package service

import (
	"github.com/br4tech/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/br4tech/meu-primeiro-crud-go/src/model"
)

func (ud *userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface) *rest_err.RestErr {
	return nil
}
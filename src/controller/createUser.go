package controller

import (
	"fmt"
	"log"

	"github.com/br4tech/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/br4tech/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/br4tech/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

 func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey","createUser"),
	)

	var UserRequest request.UserRequest

	 if err := c.ShouldBindJSON(&UserRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s\n", err.Error())
		errRest := validation.ValidateUserError(err)
		
		c.JSON(errRest.Code,errRest)
		return
	 }

	 fmt.Println(UserRequest)
 }
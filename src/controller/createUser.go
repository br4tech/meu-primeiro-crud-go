package controller

import (
	"github.com/br4tech/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

 func CreateUser(c *gin.Context) {
	 err := rest_err.NewBadRequestError("Voce chamo a rota de forma errada")
	 c.JSON(err.Code, err)
 }
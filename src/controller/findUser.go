package controller

import "github.com/gin-gonic/gin"

// gin.Context  -> representa o objeto de request, o contexto da requisicao

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {}

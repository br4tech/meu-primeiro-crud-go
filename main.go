package main

import (
	"log"

	"github.com/br4tech/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/br4tech/meu-primeiro-crud-go/src/controller"
	"github.com/br4tech/meu-primeiro-crud-go/src/controller/routes"
	"github.com/br4tech/meu-primeiro-crud-go/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// init dependencias

	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default() // instancia dois handler
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

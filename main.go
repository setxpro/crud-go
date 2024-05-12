package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/setxpro/crud-go/src/configurations/logger"
	controller "github.com/setxpro/crud-go/src/controllers"
	"github.com/setxpro/crud-go/src/controllers/routes"
	service "github.com/setxpro/crud-go/src/models/services"
)

// Ponto de entrada
func main() {

	logger.Info("About to start user applciation")

	// Carrega as variáveis de ambiente e verifica se houver alguma falha
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}

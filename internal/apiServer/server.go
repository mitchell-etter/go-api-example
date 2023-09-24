package apiServer

import (
	"log"

	"go-api-example/internal/controllers"
	"go-api-example/internal/services"

	"github.com/gin-gonic/gin"
)

type ApiServer struct {
}

func NewApiServer() *ApiServer {
	return &ApiServer{}
}

func (s *ApiServer) Run() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	// Call SetTrustedProxies to only trust specific values
	// route.SetTrustedProxies([]string{"127.0.0.1"})

	defaultController := controllers.NewDefaultController()
	defaultController.ConfigureRoutes(&router.RouterGroup)

	catalogItemService := services.NewCatalogItemService()
	skusController := controllers.NewCatalogItemsController(&catalogItemService)
	skusController.ConfigureRoutes(&router.RouterGroup)

	err := router.Run(":8080")
	if err != nil {
		log.Panicf("Starting router faled:\n%v\n", err)
	}
}

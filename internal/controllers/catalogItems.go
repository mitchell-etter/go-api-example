package controllers

import (
	"log"
	"net/http"

	"go-api-example/internal/services"

	"github.com/gin-gonic/gin"
)

type CatalogItemsController struct {
	catalogItemService services.CatalogItemService
}

func NewCatalogItemsController(catalogItemService *services.CatalogItemService) *CatalogItemsController {
	return &CatalogItemsController{
		catalogItemService: *catalogItemService,
	}
}

func (controller *CatalogItemsController) ConfigureRoutes(router *gin.RouterGroup) {
	router.GET("/catalogItems", controller.ListCatalogItems)
	router.GET("/catalogItems/:id", controller.GetCatalogItem)
}

func (controller *CatalogItemsController) ListCatalogItems(context *gin.Context) {
	items := controller.catalogItemService.GetAll()
	context.JSON(http.StatusOK, items)
}

func (controller *CatalogItemsController) GetCatalogItem(context *gin.Context) {
	sku := context.Param("id")
	catalogItem := controller.catalogItemService.GetBySku(sku)

	if catalogItem == nil {
		log.Println("catalog item not found")
		context.String(http.StatusNotFound, "")
		return
	}

	context.JSON(http.StatusOK, catalogItem)
}

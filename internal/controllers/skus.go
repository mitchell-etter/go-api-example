package controllers

import (
	"net/http"

	"go-api-example/internal/services"

	"github.com/gin-gonic/gin"
)

type SkusController struct {
	skuService services.SkuService
}

func NewSkusController(skuService *services.SkuService) *SkusController {
	return &SkusController{
		skuService: *skuService,
	}
}

func (controller *SkusController) ConfigureRoutes(router *gin.RouterGroup) {
	router.GET("/skus", controller.ListSkus)
}

func (controller *SkusController) ListSkus(context *gin.Context) {
	items := controller.skuService.GetAll()
	context.JSON(http.StatusOK, items)
}

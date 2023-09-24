package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultController struct {
}

func NewDefaultController() *DefaultController {
	return &DefaultController{}
}

func (controller *DefaultController) ConfigureRoutes(router *gin.RouterGroup) {
	router.GET("/", controller.GetRoot)
}

func (controller *DefaultController) GetRoot(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

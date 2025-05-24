package routes

import (
	"go-gin/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(routerGroup *gin.RouterGroup) {
	productController := controller.ProductController{}

	routerGroup.GET("/products", productController.GetProducts)
	routerGroup.GET("/products/:sku", productController.GetProductBySku)
}

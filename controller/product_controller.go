package controller

import (
	"go-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct{}

var productService = service.ProductService{}

func (ec *ProductController) GetProducts(c *gin.Context) {
	products := productService.GetProducts()

	c.JSON(http.StatusOK, products)
}

func (ec *ProductController) GetProductBySku(c *gin.Context) {
	sku := c.Param("sku")
	product := productService.GetProductBySku(sku)

	if product.SKU == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, product)
}

package controller

import (
	"go-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct{}

var productService = service.ProductService{}

// GetProducts godoc
// @Summary Get all products
// @Description Get all products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func (ec *ProductController) GetProducts(c *gin.Context) {
	products := productService.GetProducts()

	c.JSON(http.StatusOK, products)
}

// GetProductBySku godoc
// @Summary Get product by SKU
// @Description Get product by SKU
// @Tags Product
// @Accept json
// @Produce json
// @Param sku path string true "SKU"
// @Success 200 {object} models.Product
// @Failure 404 not found
func (ec *ProductController) GetProductBySku(c *gin.Context) {
	sku := c.Param("sku")
	product := productService.GetProductBySku(sku)

	if product.SKU == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, product)
}

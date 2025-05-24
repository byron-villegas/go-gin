package service

import (
	"go-gin/models"
	"go-gin/repository"
)

type ProductService struct{}

var productRepository = repository.ProductRepository{}

func (c ProductService) GetProducts() []models.Product {
	products := productRepository.GetProducts()

	return products
}

func (c ProductService) GetProductBySku(sku string) models.Product {
	product := productRepository.GetProductBySku(sku)

	return product
}

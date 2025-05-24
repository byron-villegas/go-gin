package repository

import (
	"encoding/json"
	"go-gin/models"
	"os"
	"strconv"
)

type ProductRepository struct{}

func (c ProductRepository) GetProducts() []models.Product {
	file, err := os.Open("data/products.json")
	if err != nil {
		return []models.Product{}
	}
	defer file.Close()

	var products []models.Product
	if err := json.NewDecoder(file).Decode(&products); err != nil {
		return []models.Product{}
	}

	return products
}

func (c ProductRepository) GetProductBySku(sku string) models.Product {
	file, err := os.Open("data/products.json")

	if err != nil {
		return models.Product{}
	}

	defer file.Close()

	var products []models.Product

	if err := json.NewDecoder(file).Decode(&products); err != nil {
		return models.Product{}
	}
	var product models.Product

	// Convert sku string to int
	skuInt, err := strconv.Atoi(sku)

	if err != nil {
		return models.Product{}
	}

	for _, p := range products {
		if p.SKU == skuInt {
			product = p
			break
		}
	}

	return product
}

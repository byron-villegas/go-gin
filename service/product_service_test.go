package service

import (
	"encoding/json"
	"go-gin/models"
	"os"
	"testing"
)

// Utilidad para crear archivo de prueba
func setupTestFile(t *testing.T, data []models.Product) func() {
	err := os.MkdirAll("data", 0755)
	if err != nil {
		t.Fatalf("failed to create data dir: %v", err)
	}
	file, err := os.Create("data/products.json")
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	if err := enc.Encode(data); err != nil {
		t.Fatalf("failed to write test data: %v", err)
	}
	return func() {
		os.Remove("data/products.json")
		os.RemoveAll("data")
	}
}

func TestGetProducts(t *testing.T) {
	testProducts := []models.Product{
		{ID: 1, Nombre: "Test Product 1"},
		{ID: 2, Nombre: "Test Product 2"},
	}
	cleanup := setupTestFile(t, testProducts)
	defer cleanup()

	service := ProductService{}
	products := service.GetProducts()

	if len(products) != len(testProducts) {
		t.Fatalf("expected %d products, got %d", len(testProducts), len(products))
	}
	for i, p := range products {
		if p.ID != testProducts[i].ID || p.Nombre != testProducts[i].Nombre {
			t.Errorf("expected product %+v, got %+v", testProducts[i], p)
		}
	}
}

func TestGetProductBySku(t *testing.T) {
	testProducts := []models.Product{
		{SKU: 1, Nombre: "Test Product 1"},
		{SKU: 2, Nombre: "Test Product 2"},
	}
	cleanup := setupTestFile(t, testProducts)
	defer cleanup()

	service := ProductService{}
	product := service.GetProductBySku("1")
	if product.SKU != 1 || product.Nombre != "Test Product 1" {
		t.Errorf("expected product %+v, got %+v", testProducts[0], product)
	}
}

// go
package repository

import (
	"encoding/json"
	"go-gin/models"
	"os"
	"testing"
)

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

func TestGetProducts_ReturnsProducts(t *testing.T) {
	testProducts := []models.Product{
		{ID: 1, Nombre: "Test Product 1"},
		{ID: 2, Nombre: "Test Product 2"},
	}
	cleanup := setupTestFile(t, testProducts)
	defer cleanup()

	repo := ProductRepository{}
	products := repo.GetProducts()

	if len(products) != len(testProducts) {
		t.Fatalf("expected %d products, got %d", len(testProducts), len(products))
	}
	for i, p := range products {
		if p.ID != testProducts[i].ID || p.Nombre != testProducts[i].Nombre {
			t.Errorf("expected product %+v, got %+v", testProducts[i], p)
		}
	}
}

func TestGetProducts_FileNotFound(t *testing.T) {
	os.Remove("data/products.json")
	os.RemoveAll("data")

	repo := ProductRepository{}
	products := repo.GetProducts()
	if len(products) != 0 {
		t.Errorf("expected empty slice, got %v", products)
	}
}

func TestGetProducts_InvalidJSON(t *testing.T) {
	err := os.MkdirAll("data", 0755)
	if err != nil {
		t.Fatalf("failed to create data dir: %v", err)
	}
	err = os.WriteFile("data/products.json", []byte("invalid json"), 0644)
	if err != nil {
		t.Fatalf("failed to write invalid json: %v", err)
	}
	defer func() {
		os.Remove("data/products.json")
		os.RemoveAll("data")
	}()

	repo := ProductRepository{}
	products := repo.GetProducts()
	if len(products) != 0 {
		t.Errorf("expected empty slice on invalid json, got %v", products)
	}
}

func TestGetProductBySku(t *testing.T) {
	testProducts := []models.Product{
		{SKU: 1, Nombre: "Test Product 1"},
		{SKU: 2, Nombre: "Test Product 2"},
	}
	cleanup := setupTestFile(t, testProducts)
	defer cleanup()

	repo := ProductRepository{}
	product := repo.GetProductBySku("1")

	if product.SKU != 1 || product.Nombre != "Test Product 1" {
		t.Errorf("expected product ID 1, got %+v", product)
	}
}

func TestGetProductBySku_NotFound(t *testing.T) {
	testProducts := []models.Product{
		{SKU: 1, Nombre: "Test Product 1"},
		{SKU: 2, Nombre: "Test Product 2"},
	}
	cleanup := setupTestFile(t, testProducts)
	defer cleanup()

	repo := ProductRepository{}
	product := repo.GetProductBySku("3")

	if product.SKU != 0 {
		t.Errorf("expected product ID 0, got %+v", product)
	}
}

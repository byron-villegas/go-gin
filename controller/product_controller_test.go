package controller

import (
	"encoding/json"
	"go-gin/models"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
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
	gin.SetMode(gin.TestMode)

	testProducts := []models.Product{
		{SKU: 1, Nombre: "Test Product 1"},
		{SKU: 2, Nombre: "Test Product 2"},
	}
	cleanup := setupTestFile(t, testProducts)
	defer cleanup()

	router := gin.Default()
	pc := &ProductController{}
	router.GET("/products", pc.GetProducts)

	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var got []models.Product
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if len(got) != len(testProducts) {
		t.Fatalf("expected %d products, got %d", len(testProducts), len(got))
	}
	for i, p := range got {
		if p.ID != testProducts[i].ID || p.Nombre != testProducts[i].Nombre {
			t.Errorf("expected product %+v, got %+v", testProducts[i], p)
		}
	}
}

func TestGetProductBySku(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testProducts := []models.Product{
		{SKU: 1, Nombre: "Test Product 1"},
		{SKU: 2, Nombre: "Test Product 2"},
	}
	cleanup := setupTestFile(t, testProducts)
	defer cleanup()

	router := gin.Default()
	pc := &ProductController{}
	router.GET("/products/:sku", pc.GetProductBySku)

	req, _ := http.NewRequest("GET", "/products/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var got models.Product
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if got.ID != testProducts[0].ID || got.Nombre != testProducts[0].Nombre {
		t.Errorf("expected product %+v, got %+v", testProducts[0], got)
	}
}

func TestGetProductBySku_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testProducts := []models.Product{
		{SKU: 1, Nombre: "Test Product 1"},
		{SKU: 2, Nombre: "Test Product 2"},
	}
	cleanup := setupTestFile(t, testProducts)
	defer cleanup()

	router := gin.Default()
	pc := &ProductController{}
	router.GET("/products/:sku", pc.GetProductBySku)

	req, _ := http.NewRequest("GET", "/products/3", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", w.Code)
	}
}

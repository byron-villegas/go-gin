package routes

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

func TestSetupRoutes_ProductsEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testProducts := []models.Product{
		{ID: 1, Nombre: "Test Product 1"},
		{ID: 2, Nombre: "Test Product 2"},
	}
	cleanup := setupTestFile(t, testProducts)
	defer cleanup()

	router := gin.Default()
	api := router.Group("/api")
	SetupRoutes(api)

	req, _ := http.NewRequest("GET", "/api/products", nil)
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

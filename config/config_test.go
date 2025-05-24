package config

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestInit_AppConfigDefaults(t *testing.T) {
	// Limpia variables de entorno relevantes
	os.Unsetenv("SECRET_KEY")
	os.Unsetenv("JWT_SECRET_KEY")
	os.Unsetenv("CORS_ORIGIN")

	// Forzamos a llamar init (en Go, init ya se ejecutó, pero AppConfig es global)
	if AppConfig == nil {
		t.Fatal("AppConfig should not be nil after init")
	}

	if AppConfig.ServerPath != "/api" {
		t.Errorf("expected ServerPath '/api', got '%s'", AppConfig.ServerPath)
	}
	if AppConfig.ServerPort != "8080" {
		t.Errorf("expected ServerPort '8080', got '%s'", AppConfig.ServerPort)
	}
	if AppConfig.SecretKey != "" {
		t.Errorf("expected SecretKey '', got '%s'", AppConfig.SecretKey)
	}
	if AppConfig.JWTSecretKey != "" {
		t.Errorf("expected JWTSecretKey '', got '%s'", AppConfig.JWTSecretKey)
	}
	if AppConfig.JWTAccessTokenExp != 3600 {
		t.Errorf("expected JWTAccessTokenExp 3600, got %d", AppConfig.JWTAccessTokenExp)
	}
	if AppConfig.CORSOrigin != "" {
		t.Errorf("expected CORSOrigin '', got '%s'", AppConfig.CORSOrigin)
	}
	if AppConfig.CORSHeaders != "Content-Type" {
		t.Errorf("expected CORSHeaders 'Content-Type', got '%s'", AppConfig.CORSHeaders)
	}
}

func TestInit_AppConfigWithEnv(t *testing.T) {
	os.Setenv("SECRET_KEY", "mysecret")
	os.Setenv("JWT_SECRET_KEY", "jwtsecret")
	os.Setenv("CORS_ORIGIN", "http://localhost")

	// Vuelve a inicializar AppConfig manualmente para el test
	AppConfig = &Config{
		ServerPath:        "/api",
		ServerPort:        "8080",
		SecretKey:         os.Getenv("SECRET_KEY"),
		JWTSecretKey:      os.Getenv("JWT_SECRET_KEY"),
		JWTAccessTokenExp: 3600,
		CORSOrigin:        os.Getenv("CORS_ORIGIN"),
		CORSHeaders:       "Content-Type",
	}

	if AppConfig.SecretKey != "mysecret" {
		t.Errorf("expected SecretKey 'mysecret', got '%s'", AppConfig.SecretKey)
	}
	if AppConfig.JWTSecretKey != "jwtsecret" {
		t.Errorf("expected JWTSecretKey 'jwtsecret', got '%s'", AppConfig.JWTSecretKey)
	}
	if AppConfig.CORSOrigin != "http://localhost" {
		t.Errorf("expected CORSOrigin 'http://localhost', got '%s'", AppConfig.CORSOrigin)
	}
}

func TestShowBanner_PrintsBannerWithPlaceholders(t *testing.T) {
	// Prepara un banner de prueba con los placeholders usados en ShowBanner
	bannerContent := `
    package.name vpackage.version
    Go: go.version
    Gin: gin.version
    Path: server.path
    Port: server.port
    `
	// Crea archivo temporal banner.txt
	err := os.WriteFile("banner.txt", []byte(bannerContent), 0644)
	if err != nil {
		t.Fatalf("failed to create banner.txt: %v", err)
	}
	defer os.Remove("banner.txt")

	// Redirige stdout para capturar la salida de ShowBanner
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	ShowBanner()

	w.Close()
	os.Stdout = stdout
	out, _ := io.ReadAll(r)

	output := string(out)

	// Verifica que los placeholders fueron reemplazados
	if !strings.Contains(output, "go-gin v1.0.0") {
		t.Errorf("expected package name and version in banner, got: %s", output)
	}
	if !strings.Contains(output, AppConfig.ServerPath) {
		t.Errorf("expected server path in banner, got: %s", output)
	}
	if !strings.Contains(output, AppConfig.ServerPort) {
		t.Errorf("expected server port in banner, got: %s", output)
	}
	if !strings.Contains(output, "Gin: "+gin.Version) {
		t.Errorf("expected gin version in banner, got: %s", output)
	}
}

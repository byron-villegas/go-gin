package config

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerPath        string
	ServerPort        string
	SecretKey         string
	JWTSecretKey      string
	JWTAccessTokenExp int
	CORSOrigin        string
	CORSHeaders       string
}

var AppConfig *Config

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: .env file not found, using default or system environment variables")
	}

	// Initialize the AppConfig struct
	AppConfig = &Config{
		ServerPath:        "/api",
		ServerPort:        "8080",
		SecretKey:         os.Getenv("SECRET_KEY"),
		JWTSecretKey:      os.Getenv("JWT_SECRET_KEY"),
		JWTAccessTokenExp: 3600,
		CORSOrigin:        os.Getenv("CORS_ORIGIN"),
		CORSHeaders:       "Content-Type",
	}
}

func ShowBanner() {
	// Get the content of the banner file
	bannerFile, err := os.ReadFile("banner.txt")
	if err != nil {
		log.Fatal("Error reading banner.txt file")
	}

	// Replace the placeholders in the banner with the actual values
	bannerLog := string(bannerFile)
	bannerLog = strings.ReplaceAll(bannerLog, "package.name", "go-gin")
	bannerLog = strings.ReplaceAll(bannerLog, "package.version", "1.0.0")
	bannerLog = strings.ReplaceAll(bannerLog, "go.version", strings.Replace(runtime.Version(), "go", "", 1))
	bannerLog = strings.ReplaceAll(bannerLog, "gin.version", gin.Version)
	bannerLog = strings.ReplaceAll(bannerLog, "server.path", AppConfig.ServerPath)
	bannerLog = strings.ReplaceAll(bannerLog, "server.port", AppConfig.ServerPort)

	// Print the banner
	fmt.Println(bannerLog)
}

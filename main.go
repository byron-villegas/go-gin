package main

import (
	"go-gin/config"
	"go-gin/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Show banner with application information
	config.ShowBanner()

	// Create a server instance
	r := gin.Default()

	// Config CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.AppConfig.CORSOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{config.AppConfig.CORSHeaders},
		AllowCredentials: true,
	}))

	// Config server path
	routerGroup := r.Group(config.AppConfig.ServerPath)

	// Config routes for the server
	routes.SetupRoutes(routerGroup)

	// Start server with port
	r.Run(":" + config.AppConfig.ServerPort)
}

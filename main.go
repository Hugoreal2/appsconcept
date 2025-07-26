package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/Hugoreal2/appsconcept/docs"
	"github.com/Hugoreal2/appsconcept/internal/handler"
	"github.com/Hugoreal2/appsconcept/internal/service"
	"github.com/Hugoreal2/appsconcept/internal/stats"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title FizzBuzz REST API
// @version 1.0
// @description A production-ready REST API server that implements a customizable FizzBuzz algorithm with request statistics tracking.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

func main() {
	// Set Gin mode based on environment
	if os.Getenv("GIN_MODE") == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize dependencies
	statsService := stats.NewService()
	fizzbuzzService := service.NewFizzBuzzService(statsService)
	fizzbuzzHandler := handler.NewFizzBuzzHandler(fizzbuzzService, statsService)

	// Setup router
	router := gin.New()

	// Add middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	v1 := router.Group("/api/v1")
	{
		v1.GET("/fizzbuzz", fizzbuzzHandler.FizzBuzz)
		v1.GET("/stats", fizzbuzzHandler.GetStats)
	}

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

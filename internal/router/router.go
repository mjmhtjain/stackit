package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mjmhtjain/stackit/internal/handlers"
)

// SetupRouter configures all routes for the application
func SetupRouter() *gin.Engine {
	r := gin.Default()

	healthHandler := handlers.NewHealthHandler()

	// Health check endpoint
	r.GET("/health", healthHandler.GetHealth)

	return r
}

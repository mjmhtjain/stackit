package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mjmhtjain/stackit/internal/service"
)

type UsageHandler struct {
	usageService service.IUsageService
}

func NewUsageHandler() *UsageHandler {
	return &UsageHandler{
		usageService: service.NewUsageService(),
	}
}

func (h *UsageHandler) PostUsage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Usage",
	})
}

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjmhtjain/stackit/internal/dto"
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
	var usageDTO dto.UsageDTO
	if err := c.ShouldBindJSON(&usageDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// Validate required fields
	if usageDTO.InstanceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "instance_id is required",
		})
		return
	}

	if usageDTO.SKUID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "skuid is required",
		})
		return
	}

	if usageDTO.Timestamp == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "timestamp is required",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Usage data received successfully",
		"data":    usageDTO,
	})
}

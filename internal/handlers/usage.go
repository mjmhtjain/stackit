package handlers

import "github.com/gin-gonic/gin"

func GetUsage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Usage",
	})
}

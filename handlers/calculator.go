package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Mosteben/calculator-api/services"
	"github.com/Mosteben/calculator-api/models"
)

func Calculate(c *gin.Context) {

	var req models.Request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON",
		})
		return
	}

	result, err := services.Calculate(req.A, req.B, req.Op)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
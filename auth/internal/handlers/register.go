package handlers

import (
	"net/http"
	"saas/auth/internal/models"

	"saas/auth/internal/services"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	input, exists := c.Get("register_input")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Missing input in context"})
		return
	}

	data := input.(models.RegisterInput)
	user, err := services.RegisterUser(data)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user.ToDTO(),
	})
}

package handlers

import (
	"fmt"
	"net/http"
	"saas/auth/internal/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	input, exists := c.Get("register_input")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Missing input in context"})
		return
	}

	data := input.(models.RegisterInput)
	fmt.Println("🔍 Register Input:", data)
	fmt.Println("📧 Email:", data.Email)
	fmt.Println("📞 Phone:", data.Phone)
	fmt.Println("🔐 Password:", data.Password)

	c.JSON(http.StatusOK, gin.H{"message": "register input received"})
}

package middlewares

import (
	"net/http"

	"saas/auth/internal/models"
	"saas/auth/internal/utils"
	"saas/auth/internal/validators"

	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateRegisterInput() gin.HandlerFunc {

	return func(c *gin.Context) {

		var input models.RegisterInput

		// Check missing fields
		if errResponse, hasError := utils.CheckMissingFields(c, &input); hasError {
			c.JSON(http.StatusBadRequest, errResponse)
			c.Abort()
			return
		}

		// Normalize and sanitize input
		email := strings.TrimSpace(strings.ToLower(input.Email))
		phone := strings.TrimSpace(input.Phone)
		password := strings.TrimSpace(input.Password)

		// Check identifier: email or phone
		if input.Email == "" && input.Phone == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email or phone number is required"})
			c.Abort()
			return
		}

		registerInput := models.RegisterInput{
			Password: password,
		}

		if email != "" {
			if !validators.IsValidEmail(email) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
				c.Abort()
				return
			}

			registerInput.Email = email
		}

		if phone != "" {
			if !validators.IsValidPhone(phone) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone number format"})
				c.Abort()
				return
			}
			registerInput.Phone = phone
		}

		// Strong password validation
		// if !strongPassRegex.MatchString(password) {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"error": "Password must be 8–64 chars, include upper, lower, number, and special character",
		// 	})
		// 	c.Abort()
		// 	return
		// }

		// Store sanitized input for handler
		c.Set("register_input", registerInput)
		c.Next()
	}
}

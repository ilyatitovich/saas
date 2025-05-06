package middlewares

import (
	"errors"
	"net/http"

	"saas/auth/internal/models"
	"saas/auth/internal/utils"
	"saas/auth/internal/validators"

	"github.com/gin-gonic/gin"
)

func ValidateRegisterInput() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.RegisterInput

		// Check for missing fields early
		if errResponse, hasError := utils.CheckMissingFields(c, &input); hasError {
			c.JSON(http.StatusBadRequest, errResponse)
			c.Abort()
			return
		}

		// Initialize sanitized input
		registerInput := models.RegisterInput{}

		if err := validateIdentifiers(&input, &registerInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if err := validatePassword(input.Password, &registerInput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Store sanitized input for the next handler
		c.Set("register_input", registerInput)
		c.Next()
	}
}

func validateIdentifiers(input *models.RegisterInput, registerInput *models.RegisterInput) error {
	if input.Email == "" && input.Phone == "" {
		return errors.New("email or phone number is required")
	}

	if input.Email != "" {
		email, err := validators.ValidateEmail(input.Email)
		if err != nil {
			return err
		}
		registerInput.Email = email
	}

	if input.Phone != "" {
		phone, err := validators.ParseAndFormatPhone(input.Phone, "US")
		if err != nil {
			return err
		}
		registerInput.Phone = phone
	}

	return nil
}

func validatePassword(password string, registerInput *models.RegisterInput) error {
	validPassword, err := validators.ValidatePassword(password)
	if err != nil {
		return err
	}
	registerInput.Password = validPassword
	return nil
}

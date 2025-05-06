package services

import (
	"errors"
	"time"

	"saas/auth/internal/models"
	"saas/auth/internal/utils"
)

// In-memory "database"
var mockUsers = []models.User{} // only for emulation

func RegisterUser(input models.RegisterInput) (models.User, error) {
	// 1. Emulate duplicate email/phone check
	for _, u := range mockUsers {
		if u.Email != "" && input.Email != "" && u.Email == input.Email {
			return models.User{}, errors.New("user with this email already exists")
		}
		if u.Phone != "" && input.Phone != "" && u.Phone == input.Phone {
			return models.User{}, errors.New("user with this phone already exists")
		}
	}

	// 2. Hash password
	hash, err := utils.HashPassword(input.Password)
	if err != nil {
		return models.User{}, err
	}

	// 3. Create user object (mock ID, timestamps)
	user := models.User{
		ID:           uint(len(mockUsers) + 1),
		Email:        input.Email,
		Phone:        input.Phone,
		PasswordHash: hash,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Role:         "user",
		Tier:         "free",
		Status:       "active",
	}

	// 4. "Insert" into mock DB
	mockUsers = append(mockUsers, user)

	// 5. Return created user (without password hash in handler ideally)
	return user, nil
}

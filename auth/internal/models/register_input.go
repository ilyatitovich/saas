package models

type RegisterInput struct {
	Email    string `json:"email" binding:"omitempty"`
	Phone    string `json:"phone" binding:"omitempty"`
	Password string `json:"password" binding:"required"`
}

package models

import (
	"time"
)

type User struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UUID string `json:"uuid" gorm:"type:uuid;default:gen_random_uuid();uniqueIndex"`

	Email        string `json:"email,omitempty" gorm:"type:varchar(255);uniqueIndex;"`
	Phone        string `json:"phone,omitempty" gorm:"type:varchar(20);uniqueIndex;"`
	PasswordHash string `json:"-" gorm:"type:varchar(255);not null"`

	FirstName string `json:"first_name,omitempty" gorm:"type:varchar(100)"`
	LastName  string `json:"last_name,omitempty" gorm:"type:varchar(100)"`

	IsVerified bool `json:"is_verified" gorm:"default:false"`
	IsActive   bool `json:"is_active" gorm:"default:true"`
	IsAdmin    bool `json:"is_admin" gorm:"default:false"`

	LastLoginAt *time.Time `json:"last_login_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"-" gorm:"index"`

	Role   string `json:"role" gorm:"type:varchar(50);default:'user'"`
	Tier   string `json:"tier" gorm:"type:varchar(50);default:'free'"`
	Status string `json:"status" gorm:"type:varchar(50);default:'active'"`
}

type UserDTO struct {
	ID    uint   `json:"id"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

func (u User) ToDTO() UserDTO {
	return UserDTO{
		ID:    u.ID,
		Email: u.Email,
		Phone: u.Phone,
	}
}

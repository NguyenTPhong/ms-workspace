package model

import "time"

type UserStatus string

const (
	StatusActive   UserStatus = "active"
	StatusInactive UserStatus = "inactive"
)

type User struct {
	Id          int64      `json:"id" gorm:"primary_key;auto_increment;not null;unique_index"`
	Email       string     `json:"email" gorm:"not null;unique_index" validate:"required,email"`
	FirstName   string     `json:"first_name" gorm:"not null" validate:"required"`
	LastName    string     `json:"last_name" gorm:"not null" validate:"required"`
	PhoneNumber string     `json:"phone_number" gorm:"not null; unique_index" validate:"required,numeric,min=10,max=15"`
	Status      UserStatus `json:"status" gorm:"not null;default:'inactive'"`
	Password    string     `json:"password" validate:"required,min=4"`
	Salt        string     `json:"salt"`
	CreatedAt   time.Time  `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
}

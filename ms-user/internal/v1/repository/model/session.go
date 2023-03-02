package model

import "time"

type Session struct {
	Id          int64      `json:"id"`
	Email       string     `json:"email"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	PhoneNumber string     `json:"phone_number"`
	Token       string     `json:"token"`
	Status      UserStatus `json:"status"`
	LoggedInAt  time.Time  `json:"logged_in_at"`
	ExpiredAt   time.Time  `json:"expired_at"`
}

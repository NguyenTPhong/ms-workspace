package model

import "time"

type EmailEvent struct {
	Event     string    `json:"event"`
	CreatedAt time.Time `json:"created_at"`
}

type Email struct {
	UserId    int64        `json:"user_id"`
	Email     string       `json:"email"`
	MessageId string       `json:"message_id"`
	Events    []EmailEvent `json:"events"`
	CreatedAt time.Time    `json:"created_at"`
}

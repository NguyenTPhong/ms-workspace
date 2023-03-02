package model

import "time"

type ActiveCode struct {
	UserId    int64     `json:"user_id" gorm:"not null"`
	Code      string    `json:"code" gorm:"not null;unique_index"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

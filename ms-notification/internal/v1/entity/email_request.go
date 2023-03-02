package entity

type SendActiveEmailRequest struct {
	UserId int64  `json:"user_id"`
	Email  string `json:"email"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Url    string `json:"url"`
}

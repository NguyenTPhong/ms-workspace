// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type HealthCheck struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type MutationHealthCheck struct {
	Text string `json:"text"`
}

type MyInfo struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Status      string `json:"status"`
}
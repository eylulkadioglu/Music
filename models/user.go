package models

type User struct {
	Email    string `json:"user_email"`
	Password string `json:"user_password"`
	Code     string `json:"code"`
}

package models

import "time"

type Program struct {
	Id   string `json:"uid"`
	Name string `json:"name"`
}

type User struct {
	Id        string    `json:"uid"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Blocked   bool      `json:"blocked"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type RegisterInput struct {
	Fullname        string
	Email           string
	Password        string
	ConfirmPassword string
}

type LoginInput struct {
	Email    string
	Password string
}

type UserPayload struct {
	JWT  string
	User User
}

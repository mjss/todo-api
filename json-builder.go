package main

import (
	"time"
)

type ErrorResponseBody struct {
	Error string `json:"error"`
}

func BuildErrorJson(e error) ErrorResponseBody {
	return ErrorResponseBody{
		Error: e.Error(),
	}
}

type JUser struct {
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func BuildUserJson(u User) JUser {
	return JUser{
		Email: u.Email,
		Name:  u.Name,
	}
}

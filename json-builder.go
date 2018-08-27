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

type JTask struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func BuildTaskJson(t Task) JTask {
	return JTask{
		Title:       t.Title,
		Description: t.Description,
		Completed:   t.Completed,
	}
}

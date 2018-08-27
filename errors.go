package main

import (
	"errors"
)

// authorization
var ErrNoAuthHeader = errors.New("no authorization header")
var ErrNoBearer = errors.New("authorization type is not bearer")

var ErrInvalidEmail = errors.New("invalid email")

var ErrPasswordTooShort = errors.New("password too short")

var ErrEmptyName = errors.New("empty name")

var ErrCreateToken = errors.New("fail to create token")

var ErrUserNotFound = errors.New("user not found")

var ErrMissingTitle = errors.New("missing title")

var ErrTitleTooLong = errors.New("Title too long")

var ErrDescriptionTooLong = errors.New("Description too long")

// db
var ErrDb = errors.New("Database returned an error")

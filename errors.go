package main

import (
	"errors"
)

var ErrInvalidEmail = errors.New("invalid email")

var ErrPasswordTooShort = errors.New("password to short")

var ErrEmptyName = errors.New("empty name")

var ErrCreateToken = errors.New("fail to create token")

var ErrUserNotFound = errors.New("user not found")

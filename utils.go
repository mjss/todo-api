package main

import (
	"errors"
	"regexp"
	"strings"
)

var ErrInvalidEmail = errors.New("invalid email")
var ErrPasswordTooShort = errors.New("password to short")
var ErrEmptyName = errors.New("empty name")
var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func isValidEmail(email string) error {
	if !emailRegexp.MatchString(email) {
		return ErrInvalidEmail
	}

	return nil
}

func isValidPassword(password string) error {
	if len(password) < 6 {
		return ErrPasswordTooShort
	}

	return nil
}

func isValidName(name string) error {
	if name == "" {
		return ErrEmptyName
	}

	return nil
}

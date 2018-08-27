package main

import (
	"regexp"
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func validateEmail(email string) error {
	if !emailRegexp.MatchString(email) {
		return ErrInvalidEmail
	}

	return nil
}

func validatePassword(password string) error {
	if len(password) < 6 {
		return ErrPasswordTooShort
	}

	return nil
}

func validateName(name string) error {
	if name == "" {
		return ErrEmptyName
	}

	return nil
}

func validateTitle(title string) error {
	if len(title) > 255 {
		return ErrTitleTooLong
	}

	if title == "" {
		return ErrMissingTitle
	}

	return nil
}

func validateDescription(description string) error {
	if len(description) > 2048 {
		return ErrDescriptionTooLong
	}

	return nil
}

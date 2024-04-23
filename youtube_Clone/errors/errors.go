package errors

import (
	"errors"
)

var (
	ErrorUserNotFound  = errors.New("users not found")
	ErrorVideoNotFound = errors.New("video not found")
)

package users

import (
	"errors"
)

var (
	// ErrUserNotFound ...
	ErrUserNotFound = errors.New("requested user could not be found")

	// ErrUserQuery ...
	ErrUserQuery = errors.New("requested users could not be retrieved base on the given criteria")

)

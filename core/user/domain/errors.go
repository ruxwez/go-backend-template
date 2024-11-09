package userDomain

import "errors"

var (
	ErrUserNotFound = errors.New("user_not_found")
	ErrToCreateUser = errors.New("error_to_create_user")
)

package authDomain

import "errors"

var (
	ErrUsernameAlreadyExists = errors.New("username_already_exists")
	ErrEmailAlreadyExists    = errors.New("email_already_exists")
	ErrInvalidCredentials    = errors.New("invalid_credentials")
	ErrInvalidToken          = errors.New("invalid_token")
	ErrTokenExpired          = errors.New("token_expired")
)

var (
	ErrEmailEmpty     = errors.New("email_empty")
	ErrUsernameEmpty  = errors.New("username_empty")
	ErrPasswordEmpty  = errors.New("password_empty")
	ErrPasswordFormat = errors.New("password_format")
	ErrEmailInvalid   = errors.New("email_invalid")
	ErrUsernameFormat = errors.New("username_format")
	ErrPasswordShort  = errors.New("password_short")
	ErrFirstNameEmpty = errors.New("first_name_empty")
	ErrLastNameEmpty  = errors.New("last_name_empty")

	// 0Auth2
	ErrInvalidState  = errors.New("invalid_state")
	ErrInvalidCode   = errors.New("invalid_code")
	ErrInvalidType   = errors.New("invalid_type")
	ErrStateNotMatch = errors.New("state_not_match")
)

var (
	ErrToVerifyTokenOnDiscord = errors.New("error_to_verify_token_on_discord")
)

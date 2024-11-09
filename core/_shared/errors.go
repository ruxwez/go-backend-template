package sharedDomain

import "errors"

var (
	ErrInvalidRequestBody = errors.New("invalid_request_body")
	ErrInternalServer     = errors.New("internal_server_error")
	ErrUnauthoried        = errors.New("unauthorized")
	ErrInvalidCookie      = errors.New("invalid_cookie")
)

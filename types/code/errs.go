package code

import "errors"

var (
	Success              = errors.New("SUCCESS")
	ErrBadRequestedData  = errors.New("ERROR_BAD_REQUESTED_DATA")
	ErrLTinyUrlNotFound  = errors.New("ERROR_TINY_URL_NOT_FOUND")
	ErrForbidden         = errors.New("ERROR_FORBIDDEN")
	ErrTinyUrlExpired    = errors.New("ERROR_TINY_URL_EXPIRED")
	ErrUserNotExists     = errors.New("ERROR_USER_NOT_EXISTS")
	ErrUserAlreadyExists = errors.New("ERROR_USER_ALREADY_EXISTS")
	ErrWrongPassword     = errors.New("ERROR_WRONG_PASSWORD")
)

package code

import "errors"

var (
	Success             = errors.New("SUCCESS")
	ErrBadRequestedData = errors.New("ERROR_BAD_REQUESTED_DATA")
	ErrLTinyUrlNotFound = errors.New("ERROR_TINY_URL_NOT_FOUND")
	ErrForbidden        = errors.New("ERROR_FORBIDDEN")
)

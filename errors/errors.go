package errors

import "errors"

var ErrNotAllowed = errors.New("not allowed")

var ErrEntityNotFound = errors.New("entity not found")

var ErrUnknown = errors.New("unknown error")

var ErrBadRequest = errors.New("bad request error")

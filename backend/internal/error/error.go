package errors

import "errors"

var ErrInvalidUrl = errors.New("invalid url format")
var ErrUnExistsUrl = errors.New("url is not exist")
var ErrUrlNotFound = errors.New("url not found")
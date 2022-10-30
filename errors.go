package http

import "errors"

var ErrInvalidRequestLine = errors.New("request line invalid")
var ErrMethodNotImplemented = errors.New("method not implemented")
var ErrInvalidHttpVersion = errors.New("invalid http version")
var ErrInvalidHeaderFormat = errors.New("invalid header format")

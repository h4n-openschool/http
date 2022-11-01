package server

import (
	"io"
	"net/http"
	"strings"
)

func NewResponse() http.Response {
	return http.Response{
		StatusCode: 200,
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     http.Header{},
	}
}

func SetBody(r http.Response, s string) http.Response {
	r.Body = io.NopCloser(strings.NewReader(s))
	return r
}

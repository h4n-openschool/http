package http

import (
	"fmt"
	"strings"
)

type StatusLine struct {
	StatusCode   int
	ReasonPhrase string
}

func (sl StatusLine) String() string {
	return fmt.Sprintf("HTTP/1.1 %v %v", sl.StatusCode, sl.ReasonPhrase)
}

type Response struct {
	StatusLine StatusLine
	Headers    map[string]string
	Body       string
}

func (r Response) String() string {
	lines := []string{}
	lines = append(lines, r.StatusLine.String())

	if len(r.Headers) > 0 {
		for k, v := range r.Headers {
			lines = append(lines, fmt.Sprintf("%v: %v", k, v))
		}
	}

	if len(r.Body) > 0 {
		lines = append(lines, "")
		lines = append(lines, strings.Join(strings.Split(r.Body, "\n"), "\r\n"))
	}

	return strings.Join(lines, "\r\n")
}

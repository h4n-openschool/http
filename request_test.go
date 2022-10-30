package http_test

import (
	"strings"
	"testing"

	"github.com/openschool/http"
	"github.com/stretchr/testify/assert"
)

func TestParseRequestLine(t *testing.T) {
	line := "GET /api/v1/classes HTTP/1.1"

	parsed, err := http.ParseRequestLine(line)

	if assert.NoError(t, err) {
		assert.Equal(t, parsed.Method, http.GET)
		assert.Equal(t, parsed.RequestUri, "/api/v1/classes")
	}
}

func TestParseHeaders(t *testing.T) {
	lines := "Accept: application/json\r\nContent-Type: application/json"

	parsed, err := http.ParseHeaders(lines)

	if assert.NoError(t, err) {
		assert.Equal(t, parsed["Accept"], "application/json")
		assert.Equal(t, parsed["Content-Type"], "application/json")
	}
}

func TestParseRequest(t *testing.T) {
	lines := []string{
		"POST /api/v1/classes HTTP/1.1",
		"Host: localhost",
		"Accept: application/json",
		"",
		`{"hello": "world"}`,
	}

	expectedBody := `{"hello": "world"}`
	expectedReqLine := http.RequestLine{
		Method:     "POST",
		RequestUri: "/api/v1/classes",
	}
	expectedHeaders := map[string]string{
		"Host":   "localhost",
		"Accept": "application/json",
	}

	parsed, err := http.ParseRequest([]byte(strings.Join(lines, "\r\n")))

	if assert.NoError(t, err) {
		assert.Equal(t, parsed.RequestLine, expectedReqLine)
		assert.Equal(t, parsed.Headers, expectedHeaders)
		assert.Equal(t, parsed.Body, expectedBody)
	}
}

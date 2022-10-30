package http_test

import (
	"testing"

	"github.com/h4n-openschool/http"
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

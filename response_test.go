package http_test

import (
	"testing"

	"github.com/openschool/http"
	"github.com/stretchr/testify/assert"
)

func TestStatusLineString(t *testing.T) {
	sl := http.StatusLine{
		StatusCode:   200,
		ReasonPhrase: "OK",
	}
	expected := "HTTP/1.1 200 OK"

	out := sl.String()

	assert.Equal(t, expected, out)
}

func TestResponseStringNoHeadersOrBody(t *testing.T) {
	response := http.Response{
		StatusLine: http.StatusLine{
			StatusCode:   200,
			ReasonPhrase: "OK",
		},
	}
	expected := "HTTP/1.1 200 OK"

	out := response.String()

	assert.Equal(t, expected, out)
}

func TestResponseStringHeadersNoBody(t *testing.T) {
	response := http.Response{
		StatusLine: http.StatusLine{
			StatusCode:   200,
			ReasonPhrase: "OK",
		},
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	expected := "HTTP/1.1 200 OK\r\nContent-Type: application/json"

	out := response.String()

	assert.Equal(t, expected, out)
}

func TestResponseStringBodyNoHeaders(t *testing.T) {
	response := http.Response{
		StatusLine: http.StatusLine{
			StatusCode:   200,
			ReasonPhrase: "OK",
		},
		Body: "hello world",
	}
	expected := "HTTP/1.1 200 OK\r\n\r\nhello world"

	out := response.String()

	assert.Equal(t, expected, out)
}

func TestResponseStringHeadersAndBody(t *testing.T) {
	response := http.Response{
		StatusLine: http.StatusLine{
			StatusCode:   200,
			ReasonPhrase: "OK",
		},
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `{"hello":"world"}`,
	}
	expected := "HTTP/1.1 200 OK\r\nContent-Type: application/json\r\n\r\n{\"hello\":\"world\"}"

	out := response.String()

	assert.Equal(t, expected, out)
}

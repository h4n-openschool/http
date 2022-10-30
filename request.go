package http

import (
	"net/textproto"
	"strings"
)

// RequestLine implements RFC2616 section 5.1.
// See https://www.rfc-editor.org/rfc/rfc2616#section-5.1
type RequestLine struct {
	// Method is one of the (supported) RFC2616 headers (defined in types.go).
	Method Method

	// RequestUri is the Request-Uri of the request received.
	RequestUri string
}

// Request represents an RFC2616 request as a struct.
type Request struct {
	// RequestLine is the parsed Request-Line of the request itself.
	RequestLine RequestLine

	// Headers represent any passed headers on the request.
	Headers textproto.MIMEHeader

	// Body optionally represents the body of the request.
	Body string
}

// ParseRequestLine takes a single line as input and returns a parsed
// Request-Line struct or an error.
func ParseRequestLine(in string) (RequestLine, error) {
	// In accordance with RFC2616, the components of the Request-Line will be
	// broken by SP characters (spaces).
	parts := strings.Split(in, " ")

	// If there are not three components of the Request-Line, it is invalid.
	if len(parts) != 3 {
		return RequestLine{}, ErrInvalidRequestLine
	}

	// If the method is not accepted by this server implementation, throw an
	// error for the client to handle.
	method, ok := methodMap[parts[0]]
	if !ok {
		return RequestLine{}, ErrMethodNotImplemented
	}

	// If the HTTP version is not 1.1, the version is not accepted by this
	// server implementation.
	if parts[2] != "HTTP/1.1" {
		return RequestLine{}, ErrInvalidHttpVersion
	}

	// Return the parsed Request-Line struct.
	return RequestLine{
		Method:     method,
		RequestUri: parts[1],
	}, nil
}

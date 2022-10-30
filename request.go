package http

import "strings"

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
	Headers map[string]string

	// Body optionally represents the body of the request.
	Body string
}

// ParseRequest parses a byte array from a TCP socket as an HTTP request.
func ParseRequest(in []byte) (Request, error) {
	// convert the byte array to a string
	data := string(in)

	// break the request into sections
	// 'sections' refer to the (request-line and headers) and the body
	sections := strings.Split(data, "\r\n\r\n") // sections will be separated by a double CRLF
	lines := strings.Split(sections[0], "\r\n") // request-line and headers will be separated by a single CRLF

	// request line should always exist
	requestLine, err := ParseRequestLine(lines[0])
	if err != nil {
		return Request{}, err
	}

	// parse the headers if there are more lines in the first section than just
	// the requsest-line
	var headers map[string]string
	if len(lines) > 1 {
		h, err := ParseHeaders(strings.Join(lines[1:], "\r\n"))
		if err != nil {
			return Request{}, err
		}

		headers = h
	}

	// if there is more than one section, parse the rest as the body
	var body string
	if len(sections) > 1 {
		// if there are more sections, assume they are the body
		body = strings.Join(sections[1:], "\r\n")
	}

	// return the parsed request struct
	return Request{
		RequestLine: requestLine,
		Headers:     headers,
		Body:        body,
	}, nil
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

// ParseHeaders will parse a collection of header lines. It takes a
// CRLF-deliminated input string and returns a map of key->value or an error if
// the header set is invalid.
func ParseHeaders(in string) (map[string]string, error) {
	headers := map[string]string{}

	// Each header will be separated by a CRLF.
	lines := strings.Split(in, "\r\n")

	// Loop over all the headers in the set and parse them
	for _, l := range lines {
		// Keys and values will be separated by a ": " string.
		kv := strings.Split(l, ": ")

		// If there is not at least a key and a value, the header is invalid
		if len(kv) < 2 {
			return map[string]string{}, ErrInvalidHeaderFormat
		}

		// Set the header value if the header is valid
		key := kv[0]
		value := strings.Join(kv[1:], ": ")
		headers[key] = value
	}

	// Return the parsed headers.
	return headers, nil
}

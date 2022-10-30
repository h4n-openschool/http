package http

type Method string

const (
	OPTIONS Method = `OPTIONS`
	GET     Method = `GET`
	HEAD    Method = `HEAD`
	POST    Method = `POST`
	PUT     Method = `PUT`
	DELETE  Method = `DELETE`
)

var methodMap = map[string]Method{
	"OPTIONS": OPTIONS,
	"GET":     GET,
	"HEAD":    HEAD,
	"POST":    POST,
	"PUT":     PUT,
	"DELETE":  DELETE,
}

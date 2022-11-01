package http

import (
	"log"
	"net"
	"net/http"
)

type HandleFunc func(req *Context) (http.Response, error)

type Context struct {
	Request    *http.Request
	RemoteAddr net.Addr
}

type Route struct {
	Method  string
	Path    string
	Handler HandleFunc
}

type Router struct {
	Routes []Route
}

func NewRouter() Router {
	return Router{}
}

func (r *Router) Route(method string, path string, handler HandleFunc) {
	route := Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}

	r.Routes = append(r.Routes, route)
}

func (r *Router) Handle() HandleFunc {
	return func(ctx *Context) (http.Response, error) {
		res := http.Response{}

		var ro Route
		found := false
		for _, route := range r.Routes {
			log.Printf("finding route: %v", route.Path)
			if ctx.Request.URL.Path == route.Path {
				log.Printf("found route: %v", route.Path)
				ro = route
				found = true
				break
			}
		}

		if !found {
			res.Status = "404 Not Found"
			return res, nil
		}

		if ctx.Request.Method != ro.Method {
			res.Status = "405 Method Not Allowed"
			return res, nil
		}

		res, err := ro.Handler(ctx)
		if err != nil {
			return res, err
		}

		return res, nil
	}
}

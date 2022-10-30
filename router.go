package http

import (
	"log"
	"net/http"
)

type HandleFunc func(req *http.Request) (http.Response, error)

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
	return func(req *http.Request) (http.Response, error) {
		res := http.Response{}

		var ro Route
		found := false
		for _, route := range r.Routes {
			log.Printf("finding route: %v", route.Path)
			if req.URL.Path == route.Path {
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

		if req.Method != ro.Method {
			res.Status = "405 Method Not Allowed"
			return res, nil
		}

		res, err := ro.Handler(req)
		if err != nil {
			return res, err
		}

		return res, nil
	}
}

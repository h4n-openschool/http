package main

import (
	"fmt"
	"log"

	"net/http"

	oshttp "github.com/h4n-openschool/http"
)

func main() {
	router := oshttp.NewRouter()
	router.Route("GET", "/hello", func(req *oshttp.Context) (http.Response, error) {
		res := oshttp.NewResponse()
		res = oshttp.SetBody(res, fmt.Sprintf("<h1>Hi %v!</h1>", req.RemoteAddr.String()))
		res.Header.Add("Content-Type", "text/html")
		res.StatusCode = 200
		return res, nil
	})

	server := oshttp.Server{
		Addr:    "127.0.0.1:8001",
		Handler: router.Handle(),
	}
	log.Fatal(server.Listen())
}

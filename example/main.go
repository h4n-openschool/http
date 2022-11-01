package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/h4n-openschool/server"
)

func main() {
	router := server.NewRouter()
	router.Route("GET", "/hello", func(req *server.Context) (http.Response, error) {
		res := server.NewResponse()
		log.Println(req.Request.URL.Query().Get("yo"))
		res = server.SetBody(res, fmt.Sprintf("<h1>Hi %v!</h1>", req.RemoteAddr.String()))
		res.Header.Add("Content-Type", "text/html")
		res.StatusCode = 200
		return res, nil
	})

	server := server.Server{
		Addr:    "127.0.0.1:8001",
		Handler: router.Handle(),
	}
	log.Fatal(server.Listen())
}

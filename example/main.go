package main

import (
	"log"

	"github.com/openschool/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8001",
	}
	log.Fatal(server.Listen())
}

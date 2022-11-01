package main

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/h4n-openschool/server"
)

var tmpl = `
<h1>Hello, world!</h1>
`

func main() {
	e := gin.Default()
	_, err := template.New("hello-world").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	e.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"hello": "world"})
	})

	server := server.Server{
		Addr:    "127.0.0.1:8001",
		Handler: e.Handler(),
	}
	log.Fatal(server.Listen())
}

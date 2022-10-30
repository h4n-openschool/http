package http

import (
	"bufio"
	"io"
	"log"
	"net"
	"net/textproto"
)

type Server struct {
	Addr string
}

func (s Server) Listen() error {
	log.Println("opening tcp socket...")
	lis, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	defer lis.Close()

	log.Println("beginning accept loop...")
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println("failed to accept connection, closing...")
			conn.Close()
		}
		log.Println("accepted connection")

		go s.Handle(conn)
	}
}

func (s Server) Handle(c net.Conn) {
	req := Request{}
	// res := Response{
	// 	StatusLine: StatusLine{
	// 		StatusCode:   200,
	// 		ReasonPhrase: "OK",
	// 	},
	// }

	reader := bufio.NewReader(c)
	tp := textproto.NewReader(reader)

	var reqLine string
	var err error
	if reqLine, err = tp.ReadLine(); err != nil {
		return
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	requestLine, err := ParseRequestLine(reqLine)
	if err != nil {
		log.Fatalf("failed to parse request line: %v", err.Error())
		return
	}
	req.RequestLine = requestLine

	header, err := tp.ReadMIMEHeader()
	if err != nil {
		log.Fatalf("failed to parse mime headers: %v", err.Error())
		return
	}
	req.Headers = header

	str, err := tp.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(str)

	c.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHi\r\n"))
	c.Close()
}

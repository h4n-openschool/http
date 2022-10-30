package http

import (
	"log"
	"net"
	"net/http"
	"net/textproto"
)

type Server struct {
	Addr    string
	Handler HandleFunc
}

func (s *Server) Listen() error {
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

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(c net.Conn) {
	defer c.Close()
	tp := textproto.NewConn(c)

	req, err1 := http.ReadRequest(tp.R)
	if err1 != nil {
		return
	}
	log.Println(req.Method)

	res, err := s.Handler(req)
	if err != nil {
		res.Status = "500 Internal Server Error"
		res.Header.Add("Content-Type", "text/plain")
	}

	err = res.Write(c)
	if err != nil {
		log.Fatalf("failed to write response: %v", err.Error())
		return
	}
	c.Close()
}

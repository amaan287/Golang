package main

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln
	<-s.quitch
	return nil
}
func (s *Server) Accept() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("Accept error: ", err)
			continue
		}
	}
}

func main() {
	log.Fatal("")
}

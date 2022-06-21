package server

import (
	"fmt"
	"goredis/server/connectionpool"
	"net"
)

type Server struct {
	Pool *connectionpool.ConnectionPool
	Host string
	Port string
}

func New() *Server {
	return &Server{
		Pool: &connectionpool.ConnectionPool{},
		Host: "127.0.0.1",
		Port: "8080",
	}
}

func (s *Server) Run() error {
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.Host, s.Port))
	if err != nil {
		return err
	}
	defer ln.Close()

	for {
		// Each new conn is in a new memory address, so we can pass it down a new channel each.
		conn, err := ln.Accept()
		if err != nil {
			return err
		}

		// Here we should handle the connection. Through the conn variable, we read and write to it.
		// What we read is the user input, and what we write is the response from the redis server.
		// So, it is reasonable to think we should handle each connection concurrently.

	}

	// We only arrive here after the server is shut down. Perhaps we look for ctrl+c signal on the server?
	return nil
}

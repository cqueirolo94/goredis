package server

import (
	"fmt"
	"goredis/server/connectionpool"
	"net"
	"time"
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
		conn, err := ln.Accept()
		if err != nil {
			return err
		}

		// Here I should get a new connection from the ConnectionPool
		go handleIncomingRequest(conn)
	}
}

func handleIncomingRequest(conn net.Conn) error {
	for {
		// store incoming data
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			return err
		}
		// respond
		time := time.Now().Format("Monday, 02-Jan-06 15:04:05 MST")
		conn.Write([]byte("Hi back!\n"))
		conn.Write([]byte(time))
	}

	// close conn
	// conn.Close()
	return nil
}

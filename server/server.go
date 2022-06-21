package server

import (
	"fmt"
	"net"
)

const max_queue_depth = 10_000_000

type Server struct {
	Host string
	Port string
}

func New() *Server {
	return &Server{
		Host: "127.0.0.1",
		Port: "8080",
	}
}

// Start the server, and all the processes that handles the connections
func (s *Server) Run() error {
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.Host, s.Port))
	if err != nil {
		return err
	}
	defer ln.Close()

	// We create the connection pool, and a channel in which we will pass the connection
	connectionChan := make(chan net.Conn, max_queue_depth)
	// Here we launch a goroutine that will be "listening" to the connectionChan, and fanning out the goroutines
	go s.handleConnections(connectionChan)

	for {
		// Each new conn is in a new memory address, so we can pass it down a new channel each.
		conn, err := ln.Accept()
		if err != nil {
			return err
		}

		// We pass the connection to the channel, that will act as a queue.
		connectionChan <- conn
	}

	// We only arrive here after the server is shut down. Perhaps we look for ctrl+c signal on the server?
	return nil
}

// Given a channel of net.Conn, we launch a goroutine to handle it
func (s *Server) handleConnections(c chan net.Conn) {
	for {
		select {
		case conn := <-c:
			// do something
		default:
			// do nothing
		}
	}
}

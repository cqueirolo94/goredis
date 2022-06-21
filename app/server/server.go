package server

import (
	"fmt"
	"net"
	"time"
)

const max_queue_depth = 8

// Models a connection, with its id and semaphore
type Connection struct {
	Sem  chan struct{}
	Conn net.Conn
	Id   int64
}

type Server struct {
	Host        string
	Port        string
	Connections chan *Connection
}

func New() *Server {
	return &Server{
		Host:        "127.0.0.1",
		Port:        "8080",
		Connections: make(chan *Connection, max_queue_depth),
	}
}

// Start the server, and all the processes that handles the connections
func (s *Server) Run() error {
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.Host, s.Port))
	if err != nil {
		return err
	}
	defer ln.Close()

	// We create a channel that will work as a semaphore.
	// If it reaches max_queue_depth all other connections will be put on hold.
	// At a new connection, we send a struct{} to it to signal a goroutine is running.
	// When a client sends the exit message, we read from this channel, freeing that space for other goroutine
	sem := make(chan struct{}, max_queue_depth)

	for {
		s.handleConnections(ln, sem)
	}

	// We only arrive here after the server is shut down. Perhaps we look for ctrl+c signal on the server?
	return nil
}

// Given a connection and a semaphore, we launch a goroutine if applicable
func (s *Server) handleConnections(ln net.Listener, sem chan struct{}) {
	select {
	case sem <- struct{}{}: // We send a struct{} to the semaphore, signaling that a goroutine will be launched.
		// We accept the connection
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}

		// Create the connection struct and send it to the channel
		s.Connections <- &Connection{
			Sem:  sem,
			Conn: conn,
			Id:   time.Now().Unix(),
		}

	default: // The semaphore was full, so we do nothing and the connection will be on hold.
		fmt.Println("Channel full. Discarding value")
	}
}

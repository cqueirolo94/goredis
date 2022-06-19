package connectionpool

import (
	"net"
	"sync"
)

// A connection pool in charge of mantain idle and active connections
type ConnectionPool struct {
	mu           sync.Mutex
	idleConns    map[string]*Connection
	numOpen      int
	maxOpenCount int
	maxIdleCount int
}

// A connection that has an id to identify it, and a channel to communicate the data back and forth
type Connection struct {
	id   string
	pool *ConnectionPool
	conn net.Conn
}

//

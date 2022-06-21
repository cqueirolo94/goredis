package main

import (
	"goredis/server"
)

func main() {
	server_ := server.New()
	server_.Run()
}

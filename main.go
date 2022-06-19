package main

import (
	"goredis/server"
)

func main() {
	//cmdmap := command.New()
	server_ := server.New()
	server_.Run()
}

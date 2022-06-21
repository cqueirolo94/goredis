package main

import (
	"goredis/app"
)

func main() {
	goredis := app.InitializeApp()
	goredis.Run()
}

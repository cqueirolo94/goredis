package app

import (
	"fmt"
	"goredis/app/command"
	"goredis/app/server"
)

const (
	buffer_size = 4000
)

// The app will manage all of the goredis blocks, like the command list and the server.
type App struct {
	srv    *server.Server
	cmdMap *command.CommandMap
}

// Creates the app struct, together with all of the modules that correspond.
func InitializeApp() *App {
	return &App{
		srv:    server.New(),
		cmdMap: command.New(),
	}
}

// Starts the app
func (app *App) Run() {
	// Start the server
	go app.srv.Run()

	// Manage openned connections
	for {
		select {
		case conn := <-app.srv.Connections:
			app.processConnection(conn)
		default:
		}
	}

}

// On each opened connection, we manage the flow: Reading from buffer, processing the command, and return the result.
func (app *App) processConnection(conn *server.Connection) {
	for {
		buff := make([]byte, buffer_size)
		_, err := conn.Conn.Read(buff)
		if err != nil {
			fmt.Println(err)
		}

		res, err := app.cmdMap.Run(buff)
		if err != nil {
			if err.Error() == "exit" {
				<-conn.Sem
				conn.Conn.Close()
				return
			}
			fmt.Println(err)
		}

		conn.Conn.Write([]byte(res + "\n"))
	}
}

func (app *App) handleError(err string) string {
	errPrefix := "Error: "
	return errPrefix
}

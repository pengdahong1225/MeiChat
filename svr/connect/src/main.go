package main

import (
	"connect/src/server"
)

func main() {
	// tcp
	cli_ := &server.Cli{}
	cli_.Run()

	// websocket
	svr_ := &server.Server{}
	svr_.Run()
}

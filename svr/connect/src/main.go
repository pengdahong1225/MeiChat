package main

import (
	"connect/src/server/connect"
	"connect/src/server/wsconnect"
)

func main() {
	// tcp
	cli_ := &connect.Cli{}
	cli_.Run()

	// websocket
	svr_ := &wsconnect.Server{}
	svr_.Run()
}

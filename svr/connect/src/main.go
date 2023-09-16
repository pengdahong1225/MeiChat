package main

import (
	"connect/src/server/tcpconnect"
	"connect/src/server/wsconnect"
)

func main() {
	// tcp
	cli_ := &tcpconnect.Cli{}
	cli_.Run()

	// websocket
	svr_ := &wsconnect.Server{}
	svr_.Run()
}

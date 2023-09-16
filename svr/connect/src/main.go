package main

import (
	"connect/src/communicator"
	"connect/src/webServer"
)

func main() {
	communicator.Start()
	svr_ := &webServer.Server{}
	svr_.Run()
}

package main

import (
	"connect/src/common/initializer"
	"connect/src/communicator"
	"connect/src/webServer"
)

func main() {
	initializer.InitServer()
	communicator.Start()
	svr_ := &webServer.Server{}
	svr_.Run()
}

package main

import (
	_ "user/src/common/initializer"
	"user/src/server"
)

func main() {
	tcpServer := server.GetInstance()
	listener := tcpServer.Start()
	tcpServer.Loop(listener)
}

package server

import (
	"fmt"
	"net"
	"user/src/common"
)

type tcpServer struct {
	connections []*net.TCPConn
}

var instance *tcpServer

func init() {
	instance = new(tcpServer)
	instance.connections = make([]*net.TCPConn, 5)
}

// factory
func GetInstance() *tcpServer {
	return instance
}

func (receiver tcpServer) Start() *net.TCPListener {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":9000")
	listener, _ := net.ListenTCP("tcp", tcpAddr)
	return listener
}

func (receiver tcpServer) Loop(listener *net.TCPListener) {
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("new connection,addr =", conn.RemoteAddr().String())
		receiver.connections = append(receiver.connections, conn)
		// 异步处理连接
		if e := common.AntsPoolInstance.Submit(func() {
			receiver.newConnectionHandle(conn)
		}); e != nil {
			fmt.Println(e)
		}
	}
}

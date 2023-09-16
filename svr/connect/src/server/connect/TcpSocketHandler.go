package connect

import (
	"connect/src/common"
	codec2 "connect/src/common/codec"
	"connect/src/common/initializer"
	pb "connect/src/proto"
	"fmt"
	"log"
	"net"
	"strconv"
)

type TcpSocketHandler struct {
	conn net.Conn
}

var SvrMap map[pb.ENPositionType]*TcpSocketHandler

func newTcpHandler(ip string, port int) *TcpSocketHandler {
	tcpHandler := TcpSocketHandler{}
	var err error
	tcpHandler.conn, err = net.Dial("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &tcpHandler
}

func (receiver TcpSocketHandler) GetConnection() net.Conn {
	return receiver.conn
}
func (receiver TcpSocketHandler) start() {
	codec := codec2.GetCodec()
	// 异步
	err := common.AntsPoolInstance.Submit(func() {
		buffer := make([]byte, 1024)
		for {
			n, err := receiver.conn.Read(buffer) // 阻塞至inner返回消息
			if err != nil {
				fmt.Println(err)
				break
			}
			frame := make([]byte, len(buffer))
			copy(frame, buffer[:n])
			// 数据包解析
			body, errDec := codec.DeCodeData(frame)
			if errDec != nil {
				log.Println(errDec)
				break
			}
			// 同步
			receiver.handleSync(body)
		}
	})
	if err != nil {
		return
	}
}

/////////////////////////////////////////////////////////////////////////////////////////////

type Cli struct {
}

func (receiver Cli) Run() {
	SvrMap = make(map[pb.ENPositionType]*TcpSocketHandler)
	if SvrMap[pb.ENPositionType_EN_Position_User] = newTcpHandler(initializer.ServerInfoInstance.User_.Host,
		initializer.ServerInfoInstance.User_.Port); SvrMap[pb.ENPositionType_EN_Position_User] == nil {
		fmt.Println("connect to user error")
		return
	}
	if SvrMap[pb.ENPositionType_EN_Position_ChatServer] = newTcpHandler(initializer.ServerInfoInstance.Chat_.Host,
		initializer.ServerInfoInstance.Chat_.Port); SvrMap[pb.ENPositionType_EN_Position_ChatServer] == nil {
		fmt.Println("connect to chatServer error")
		return
	}

	// 启动客户端 -- 连接内部服务
	for _, svr := range SvrMap {
		svr.start()
	}
}

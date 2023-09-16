package tcpconnect

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

type InnerCommunicator struct {
}

func (receiver InnerCommunicator) Run() {
	common.SvrMap = make(map[pb.ENPositionType]*net.Conn)
	if common.SvrMap[pb.ENPositionType_EN_Position_User] = newTcpHandler(initializer.ServerInfoInstance.User_.Host,
		initializer.ServerInfoInstance.User_.Port); common.SvrMap[pb.ENPositionType_EN_Position_User] == nil {
		fmt.Println("tcpconnect to user error")
		return
	}
	if common.SvrMap[pb.ENPositionType_EN_Position_ChatServer] = newTcpHandler(initializer.ServerInfoInstance.Chat_.Host,
		initializer.ServerInfoInstance.Chat_.Port); common.SvrMap[pb.ENPositionType_EN_Position_ChatServer] == nil {
		fmt.Println("tcpconnect to chatServer error")
		return
	}

	// 启动客户端 -- 连接内部服务
	for _, conn := range common.SvrMap {
		loop(*conn)
	}
}

func loop(conn net.Conn) {
	codec := codec2.GetCodec()
	// 异步
	err := common.AntsPoolInstance.Submit(func() {
		buffer := make([]byte, 1024)
		for {
			n, err := conn.Read(buffer) // 阻塞至inner返回消息
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

func newTcpHandler(ip string, port int) *net.Conn {
	conn, err := net.Dial("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &conn
}

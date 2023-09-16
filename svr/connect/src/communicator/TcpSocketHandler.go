package communicator

import (
	"connect/src/common"
	codec2 "connect/src/common/codec"
	"fmt"
	"log"
	"net"
)

func Start() {
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
			handleSync(body)
		}
	})
	if err != nil {
		return
	}
}

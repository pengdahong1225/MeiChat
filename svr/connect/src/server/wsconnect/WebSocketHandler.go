package wsconnect

import (
	"connect/src/common"
	codec2 "connect/src/common/codec"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin:     func(r *http.Request) bool { return true },
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	ConnectionsMap map[int64]*websocket.Conn // (uid,客户端连接)
)

type Server struct {
}

func (receiver Server) Run() {
	ConnectionsMap = make(map[int64]*websocket.Conn)
	// 启动
	http.HandleFunc("/ws", receiver.websocketHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (receiver Server) websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // 将http升级到ws连接
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	fmt.Println("new connection,addr =", conn.RemoteAddr().String())

	codec := codec2.GetCodec()

	wgHandle := new(sync.WaitGroup)
	wgHandle.Add(1)
	err = common.AntsPoolInstance.Submit(func() {
		defer wgHandle.Done()
		for {
			// 读取客户端发送的消息
			_, data, errRead := conn.ReadMessage()
			if errRead != nil {
				log.Println(err)
				break
			}
			// 数据包解析
			body, errDec := codec.DeCodeData(data)
			if errDec != nil {
				log.Println(errDec)
				break
			}
			// 异步
			bodyTmp := make([]byte, len(body))
			copy(bodyTmp, body)
			receiver.handleAsync(bodyTmp, conn)
		}
	})
	if err != nil {
		return
	}
	wgHandle.Wait()
}
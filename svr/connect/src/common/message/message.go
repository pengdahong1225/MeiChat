package message

import (
	codec2 "connect/src/common/codec"
	pb "connect/src/proto"
	"connect/src/server"
	"fmt"
	"github.com/gorilla/websocket"
)

type Message struct {
	WebSocketHandler *websocket.Conn
	SocketHandler    *server.TcpSocketHandler
}

func (receiver Message) SendRequestToUser(head *pb.PBHead, msg *pb.PBCMsg) {
	codec := codec2.GetCodec()
	data, err := codec.EnCodeMsg(head, msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	frame, e := codec.EnCodeData(data)
	if e != nil {
		fmt.Println(err)
		return
	}
	_, ew := receiver.SocketHandler.GetConn().Write(frame)
	if ew != nil {
		fmt.Println(err)
		return
	}
}

func (receiver Message) SendResponseToClient(head *pb.PBHead, msg *pb.PBCMsg) bool {
	codec := codec2.GetCodec()
	data, err := codec.EnCodeMsg(head, msg)
	if err != nil {
		fmt.Println(err)
		return false
	}
	frame, e := codec.EnCodeData(data)
	if e != nil {
		fmt.Println(err)
		return false
	}
	if receiver.WebSocketHandler.WriteMessage(websocket.BinaryMessage, frame) == nil {
		return true
	}
	return false
}

func (receiver Message) SendRequestToChatServer(head *pb.PBHead, msg *pb.PBCMsg) {
	codec := codec2.GetCodec()
	data, err := codec.EnCodeMsg(head, msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	frame, e := codec.EnCodeData(data)
	if e != nil {
		fmt.Println(err)
		return
	}
	_, _ = receiver.SocketHandler.GetConn().Write(frame)
}

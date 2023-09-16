package message

import (
	codec2 "connect/src/common/codec"
	pb "connect/src/proto"
	"fmt"
	"github.com/gorilla/websocket"
	"net"
)

func SendRequestToUser(conn net.Conn, head *pb.PBHead, msg *pb.PBCMsg) {
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
	_, ew := conn.Write(frame)
	if ew != nil {
		fmt.Println(err)
		return
	}
}

func SendResponseToClient(conn *websocket.Conn, head *pb.PBHead, msg *pb.PBCMsg) bool {
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
	if conn.WriteMessage(websocket.BinaryMessage, frame) == nil {
		return true
	}
	return false
}

func SendRequestToChatServer(conn net.Conn, head *pb.PBHead, msg *pb.PBCMsg) {
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
	_, _ = conn.Write(frame)
}

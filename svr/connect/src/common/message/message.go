package message

import (
	"connect/src/common"
	codec2 "connect/src/common/codec"
	pb "connect/src/proto"
	"github.com/gorilla/websocket"
	"log"
)

func SendRequestToUser(head *pb.PBHead, msg *pb.PBCMsg) {
	conn := *common.SvrMap[head.Route.Destination]

	codec := codec2.GetCodec()
	data, err := codec.EnCodeMsg(head, msg)
	if err != nil {
		log.Println(err)
		return
	}
	frame, e := codec.EnCodeData(data)
	if e != nil {
		log.Println(err)
		return
	}
	_, ew := conn.Write(frame)
	if ew != nil {
		log.Println(err)
		return
	}
}

func SendRequestToChatServer(head *pb.PBHead, msg *pb.PBCMsg) {
	conn := *common.SvrMap[head.Route.Destination]

	codec := codec2.GetCodec()
	data, err := codec.EnCodeMsg(head, msg)
	if err != nil {
		log.Println(err)
		return
	}
	frame, e := codec.EnCodeData(data)
	if e != nil {
		log.Println(err)
		return
	}
	_, _ = conn.Write(frame)
}

func SendResponseToClient(head *pb.PBHead, msg *pb.PBCMsg) bool {
	websocketHandler := common.ConnectionsMap[head.Uid]

	codec := codec2.GetCodec()
	data, err := codec.EnCodeMsg(head, msg)
	if err != nil {
		log.Println(err)
		return false
	}
	frame, e := codec.EnCodeData(data)
	if e != nil {
		log.Println(err)
		return false
	}
	if websocketHandler.WriteMessage(websocket.BinaryMessage, frame) == nil {
		return true
	}
	return false
}

func SendMsgToClient(head *pb.PBHead, msg *pb.PBCMsg) bool {
	websocketHandler := common.ConnectionsMap[msg.GetCsResponseChatSingle().DstUid]
	codec := codec2.GetCodec()
	data, err := codec.EnCodeMsg(head, msg)
	if err != nil {
		log.Println(err)
		return false
	}
	frame, e := codec.EnCodeData(data)
	if e != nil {
		log.Println(err)
		return false
	}
	if websocketHandler.WriteMessage(websocket.BinaryMessage, frame) == nil {
		return true
	}
	return false
}

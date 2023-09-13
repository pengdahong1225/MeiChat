package message

import (
	"github.com/gorilla/websocket"
	"user/src/common/session"
	pb "user/src/proto"
)

type Message struct {
	WebSocketHandler *websocket.Conn
	//SocketHandler    *server.TcpSocketHandler
}

func (receiver Message) SendRequestToUser(psession *session.Session, head *pb.PBHead, msg *pb.PBCMsg) {

}

func (receiver Message) SendRequestToClient(psession *session.Session, head *pb.PBHead, msg *pb.PBCMsg) {

}

func (receiver Message) Send(head *pb.PBHead, msg *pb.PBCMsg) {

}

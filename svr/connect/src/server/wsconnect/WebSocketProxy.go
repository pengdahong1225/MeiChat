package wsconnect

import (
	"connect/src/common"
	codec2 "connect/src/common/codec"
	"connect/src/common/message"
	"connect/src/common/session"
	"connect/src/processer"
	pb "connect/src/proto"
	"connect/src/server/connect"
	"github.com/gorilla/websocket"
)

// 异步处理
func (receiver Server) handleAsync(body []byte, conn *websocket.Conn) {
	err := common.AntsPoolInstance.Submit(func() {
		receiver.handleSync(body, conn)
	})
	if err != nil {
		return
	}
}

// 同步处理
func (receiver Server) handleSync(body []byte, conn *websocket.Conn) {
	// 解包
	codec := codec2.GetCodec()
	header, msg, err := codec.DeCodeMsg(body)
	if err != nil {
		println(err)
		return
	}
	// 绑定uid和conn
	ConnectionsMap[header.Uid] = conn

	// 处理
	receiver.handle(header, msg)
	return
}

// 用于处理从客户端来的包
func (receiver Server) handle(head *pb.PBHead, msg *pb.PBCMsg) {
	var s *session.Session
	// 获取session
	s = session.ManagerInstance.GetSession(int(head.SessionId))
	if s == nil {
		// 分配新的session [客户端来的包一般都是新请求，可以直接分配]
		s = session.ManagerInstance.AllocSession()
	}
	s.Head_ = head
	s.RequestMsg_ = msg
	s.Head_.SessionId = int32(s.SessionID) // 更新id
	s.SessionState_ = session.EN_Session_Idle
	s.MessageType_ = s.Head_.Mtype

	if !isProcessInLocal(msg) {
		doTransfer(head, msg)
	} else {
		processer.Instance().Process(s)
	}
}

func isProcessInLocal(msg *pb.PBCMsg) bool {
	switch msg.MsgUnion.(type) {
	case *pb.PBCMsg_CsRequestLogin:
		return true
	case *pb.PBCMsg_CsRequestRegist:
		return true
	}
	return false
}

func doTransfer(head *pb.PBHead, msg *pb.PBCMsg) {
	socketHandler_ := connect.SvrMap[head.Route.Destination]
	message.SendRequestToChatServer(socketHandler_.GetConnection(), head, msg)
}

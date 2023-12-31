package communicator

import (
	codec2 "connect/src/common/codec"
	"connect/src/common/session"
	"connect/src/processer"
	pb "connect/src/proto"
	"log"
)

func handleSync(data []byte) {
	// 解包
	codec := codec2.GetCodec()
	header, msg, err := codec.DeCodeMsg(data)
	if err != nil {
		println(err)
		return
	}
	// 处理
	handle(header, msg)
}

// 用于处理从其他服返回的包(user,chatServer)
func handle(head *pb.PBHead, msg *pb.PBCMsg) {
	// 查找session
	sessionID := head.SessionId
	psession := session.ManagerInstance.GetSession(int(sessionID))
	if psession == nil {
		// 错误 [如果session有定时器，可能是超时返回，否则就是error]
		log.Printf("failed to found session[%d]\n", sessionID)
	} else {
		if psession.Head_.Uid != head.Uid {
			log.Printf("session can't match uid\n")
			return
		}
		psession.MessageType_ = head.Mtype
		psession.ResponseMsg_ = msg

		processer.Instance().Process(psession)
	}
}

package processer

import (
	"connect/src/common/message"
	"connect/src/common/session"
	"connect/src/processer/internal"
	pb "connect/src/proto"
)

type processerManager struct {
}

var instance *processerManager

func init() {
	if instance == nil {
		instance = new(processerManager)
	}
}

func Instance() *processerManager {
	return instance
}

func (receiver processerManager) Process(psession *session.Session) {
	if !isProcessInLocal(psession.RequestMsg_) {
		message.SendRequestToChatServer(psession.Head_, psession.RequestMsg_)
	} else {
		internal.Do(psession)
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

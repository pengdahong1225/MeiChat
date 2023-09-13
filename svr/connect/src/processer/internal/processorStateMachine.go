package internal

import (
	"connect/src/common/session"
	pb "connect/src/proto"
	"fmt"
)

// ENHandlerResult
const (
	EN_Handler_Done = iota
	EN_Handler_Succ
)

// 状态机 -- 多态处理请求
func ProcessEvent(process_ process) {
	result := EN_Handler_Done
	if process_.GetSession().MessageType_ == pb.ENMessageType_EN_Message_Request {
		process_.GetSession().SessionState_ = session.EN_Session_Idle
		result = process_.ProcessRequestMsg()
	} else if process_.GetSession().MessageType_ == pb.ENMessageType_EN_Message_Response {
		if process_.GetSession().SessionState_ == session.EN_Session_Idle {
			result = process_.ProcessResponseMsg()
		}
	}
	switch result {
	case EN_Handler_Done:
		EndProcess(process_.GetSession()) // 结束回话，释放session
		return
	case EN_Handler_Succ:
		return
	default:
		return
	}
}

func EndProcess(psession *session.Session) {
	session.ManagerInstance.ReleaseSession(psession.SessionID)
	fmt.Printf("session[%d] is release\n", psession.SessionID)
}

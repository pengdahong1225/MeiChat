package internal

import (
	"user/src/common/session"
	pb "user/src/proto"
)

// 用户数据注册服务
type addDataRequest struct {
	session_ *session.Session
}

func NewAddDataRequest() *addDataRequest {
	return &addDataRequest{
		session_: nil,
	}
}

func (receiver addDataRequest) SetSession(psession *session.Session) {
	receiver.session_ = psession
}
func (receiver addDataRequest) GetSession() *session.Session {
	return receiver.session_
}
func (receiver addDataRequest) ProcessRequestMsg() *pb.PBCMsg {
	return nil
}

package internal

import (
	"user/src/common/session"
	pb "user/src/proto"
)

// 用户数据注册服务
type registRequest struct {
	session_ *session.Session
}

func NewRegistRequest() *registRequest {
	return &registRequest{
		session_: nil,
	}
}

func (receiver registRequest) SetSession(psession *session.Session) {
	receiver.session_ = psession
}
func (receiver registRequest) GetSession() *session.Session {
	return receiver.session_
}
func (receiver registRequest) ProcessRequestMsg() *pb.PBCMsg {

	return nil
}

package internal

import (
	"user/src/common/session"
	pb "user/src/proto"
)

type pushDataRequest struct {
	session_ *session.Session
}

func NewPushDataRequest() *pushDataRequest {
	return &pushDataRequest{
		session_: nil,
	}
}

func (receiver pushDataRequest) SetSession(psession *session.Session) {
	receiver.session_ = psession
}
func (receiver pushDataRequest) GetSession() *session.Session {
	return receiver.session_
}
func (receiver pushDataRequest) ProcessRequestMsg() *pb.PBCMsg {
	return nil
}

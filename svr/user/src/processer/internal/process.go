package internal

import (
	"user/src/common/session"
	pb "user/src/proto"
)

// process总接口
type process interface {
	SetSession(psession *session.Session)
	GetSession() *session.Session
	ProcessRequestMsg() *pb.PBCMsg
}

var _handler_map map[pb.PBCMsgCmd]process

func init() {
	_handler_map = make(map[pb.PBCMsgCmd]process)
	_handler_map[pb.PBCMsgCmd_ss_request_login] = NewLoginRequest()
	_handler_map[pb.PBCMsgCmd_ss_request_regist] = NewRegistRequest()
}

func getProcesser(cmd_ pb.PBCMsgCmd) process {
	for cmd, p := range _handler_map {
		if cmd == cmd_ {
			return p
		}
	}
	return nil
}

func Do(psession *session.Session) *pb.PBCMsg {
	p := getProcesser(psession.Head_.Cmd)
	p.SetSession(psession)
	return p.ProcessRequestMsg()
}

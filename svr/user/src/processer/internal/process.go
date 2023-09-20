package internal

import (
	"user/src/common/session"
	pb "user/src/proto"
)

// cmd
const (
	cs_request_login = iota + 1
	cs_response_login
	cs_request_regist
	cs_response_regist
	cs_request_chat_single
	cs_response_chat_single
	cs_request_chat_group
	cs_response_chat_group
	ss_request_get_user_data  = 10
	ss_response_get_user_data = iota + 2
	ss_request_push_data
	ss_response_push_data
	ss_request_add_data
	ss_response_add_data
)

// process总接口
type process interface {
	SetSession(psession *session.Session)
	GetSession() *session.Session
	ProcessRequestMsg() *pb.PBCMsg
}

var _handler_map map[int32]process

func init() {
	_handler_map = make(map[int32]process)
	_handler_map[ss_request_get_user_data] = NewGetDataRequest()
	_handler_map[ss_request_add_data] = NewAddDataRequest()
	_handler_map[ss_request_push_data] = NewPushDataRequest()
}

func getProcesser(cmd_ int32) process {
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

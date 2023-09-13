package internal

import (
	"connect/src/common/session"
	pb "connect/src/proto"
)

const (
	EN_Purpose_Get_Data = iota
	EN_Purpose_Update_Data
)

// process总接口
type process interface {
	SetSession(psession *session.Session)
	GetSession() *session.Session
	ProcessRequestMsg() int
	ProcessResponseMsg() int
}

// TODO 在多线程可能存在资源竞争
var _handler_map map[pb.PBCMsgCmd]process

func init() {
	_handler_map = make(map[pb.PBCMsgCmd]process)
	_handler_map[pb.PBCMsgCmd_cs_request_login] = newLoginProcesser()
	_handler_map[pb.PBCMsgCmd_cs_request_regist] = newRegistProcesser()
	_handler_map[pb.PBCMsgCmd_cs_response_chat_single] = newChatSingleProcesser()
	_handler_map[pb.PBCMsgCmd_cs_response_chat_group] = newChatGroupProcesser()
}

func getProcesser(cmd_ pb.PBCMsgCmd) process {
	for cmd, p := range _handler_map {
		if cmd == cmd_ {
			return p
		}
	}
	return nil
}

// internal 入口
func Do(psession *session.Session) {
	p := getProcesser(psession.Head_.Cmd)
	p.SetSession(psession) // 设置session
	ProcessEvent(p)
}

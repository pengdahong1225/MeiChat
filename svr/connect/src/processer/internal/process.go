package internal

import (
	"connect/src/common/session"
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
	ProcessRequestMsg() int
	ProcessResponseMsg() int
}

// TODO 在多线程可能存在资源竞争
var _handler_map map[int32]process

func init() {
	_handler_map = make(map[int32]process)
	_handler_map[cs_request_login] = newLoginProcesser()
	_handler_map[cs_request_regist] = newRegistProcesser()
	_handler_map[cs_request_chat_single] = newChatSingleProcesser()
	_handler_map[cs_request_chat_group] = newChatGroupProcesser()
}

func getProcesser(cmd_ int32) process {
	return _handler_map[cmd_]
}

// internal入口
func Do(psession *session.Session) {
	p := getProcesser(psession.Head_.Cmd)
	p.SetSession(psession) // 设置session
	ProcessEvent(p)
}

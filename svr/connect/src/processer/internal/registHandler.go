package internal

import (
	"connect/src/common/session"
)

// 注册请求
type registProcesser struct {
	psession *session.Session
}

func newRegistProcesser() *registProcesser {
	return &registProcesser{}
}

func (receiver registProcesser) SetSession(p *session.Session) {
	receiver.psession = p
}

func (receiver registProcesser) GetSession() *session.Session {
	return receiver.psession
}

func (receiver registProcesser) ProcessRequestMsg() int {
	return -1
}

func (receiver registProcesser) ProcessResponseMsg() int {
	return -1
}

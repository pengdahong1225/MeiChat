package internal

import (
	"connect/src/common/session"
)

type chatGroupProcesser struct {
	psession *session.Session
}

func newChatGroupProcesser() *chatSingleProcesser {
	return &chatSingleProcesser{}
}

func (receiver chatGroupProcesser) SetSession(p *session.Session) {
	receiver.psession = p
}

func (receiver chatGroupProcesser) GetSession() *session.Session {
	return receiver.psession
}

func (receiver chatGroupProcesser) ProcessRequestMsg() int {
	return EN_Handler_Done
}

func (receiver chatGroupProcesser) ProcessResponseMsg() int {
	return EN_Handler_Done
}

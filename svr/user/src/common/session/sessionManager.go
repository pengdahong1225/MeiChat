package session

import (
	"fmt"
	"user/src/common"
)

type Manager struct {
	sessionMap map[int]*Session
	freeQue    common.Queue[int]
}

var (
	ManagerInstance *Manager
)

func init() {
	ManagerInstance = &Manager{}
	ManagerInstance.sessionMap = make(map[int]*Session)
	for i := 0; i < 100; i++ {
		ManagerInstance.freeQue.Push(i)
	}
}

func (receiver *Manager) AllocSession() *Session {
	sessionID := allocSessionID()
	if sessionID < 0 {
		return nil
	}
	s := newSession(sessionID)
	ManagerInstance.sessionMap[sessionID] = s
	return s
}

func (receiver *Manager) GetSession(sessionID int) *Session {
	if s, exist := ManagerInstance.sessionMap[sessionID]; exist {
		return s
	}
	return nil
}

func (receiver *Manager) ReleaseSession(sessionID int) {
	if _, exist := ManagerInstance.sessionMap[sessionID]; exist {
		delete(ManagerInstance.sessionMap, sessionID)
		ManagerInstance.freeQue.Push(sessionID)
	}
}

func allocSessionID() int {
	if ManagerInstance.freeQue.GetLength() <= 0 {
		return -1
	}
	var freeID int
	err := ManagerInstance.freeQue.Pop(&freeID)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return freeID
}

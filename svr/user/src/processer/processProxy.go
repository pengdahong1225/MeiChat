package processer

import (
	"user/src/common/session"
	"user/src/processer/internal"
	pb "user/src/proto"
)

type processerManager struct {
}

var instance *processerManager

func init() {
	if instance == nil {
		instance = new(processerManager)
	}
}

func Instance() *processerManager {
	return instance
}

func (receiver processerManager) Process(psession *session.Session) *pb.PBCMsg {
	psession.MessageType_ = psession.Head_.Route.Mtype
	return internal.Do(psession)
}

package session

import pb "connect/src/proto"

// ENSessionState
const (
	EN_Session_Idle = iota
	EN_Session_Wait_Get_Data
	EN_Session_Wait_Update_Data
)

type Session struct {
	SessionID     int
	Head_         *pb.PBHead
	RequestMsg_   *pb.PBCMsg
	ResponseMsg_  *pb.PBCMsg
	SessionState_ int
	MessageType_  pb.ENMessageType
}

// factory
func newSession(id int) *Session {
	return &Session{
		SessionID:     id,
		Head_:         nil,
		RequestMsg_:   nil,
		ResponseMsg_:  nil,
		SessionState_: -1,
		MessageType_:  -1,
	}
}

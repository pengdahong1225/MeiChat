package internal

import (
	"connect/src/common/message"
	"connect/src/common/session"
	pb "connect/src/proto"
)

type chatSingleProcesser struct {
	psession *session.Session
}

func newChatSingleProcesser() *chatSingleProcesser {
	return &chatSingleProcesser{}
}

func (receiver chatSingleProcesser) SetSession(p *session.Session) {
	receiver.psession = p
}

func (receiver chatSingleProcesser) GetSession() *session.Session {
	return receiver.psession
}

func (receiver chatSingleProcesser) ProcessRequestMsg() int {
	return EN_Handler_Done
}

// 从chatServer返回的包 -- 根据result发送给client
func (receiver chatSingleProcesser) ProcessResponseMsg() int {
	ss_response := receiver.psession.ResponseMsg_.GetCsResponseChatSingle()
	var msg *pb.PBCMsg
	response := *msg.GetCsResponseChatSingle()
	response.Result = ss_response.Result
	route := &pb.PBRoute{
		Source:      pb.ENPositionType_EN_Position_Connect,
		Destination: pb.ENPositionType_EN_Position_Client,

		RouteType: pb.ENRouteType_EN_Route_p2p,
	}
	head := &pb.PBHead{
		Route:     route,
		Uid:       receiver.psession.Head_.Uid,
		Cmd:       cs_response_chat_single,
		SessionId: int32(receiver.psession.SessionID),
		Mtype:     pb.ENMessageType_EN_Message_Response,
	}

	// response
	message.SendResponseToClient(head, msg)

	// 通知dst客户端
	if response.Result == pb.ENMessageError_EN_MESSAGE_ERROR_OK {
		message.SendMsgToClient(head, msg)
	}
	// TODO log

	return EN_Handler_Done
}

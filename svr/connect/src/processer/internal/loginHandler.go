package internal

import (
	"connect/src/common/message"
	"connect/src/common/session"
	pb "connect/src/proto"
	"connect/src/server"
)

// 登录请求
type loginProcesser struct {
	psession *session.Session
}

func newLoginProcesser() *loginProcesser {
	return &loginProcesser{}
}

func (receiver loginProcesser) SetSession(p *session.Session) {
	receiver.psession = p
}

func (receiver loginProcesser) GetSession() *session.Session {
	return receiver.psession
}

func (receiver loginProcesser) ProcessRequestMsg() int {
	// 拉取用户数据
	route := &pb.PBRoute{
		Source:      pb.ENPositionType_EN_Position_Connect,
		Destination: pb.ENPositionType_EN_Position_User,
		SessionId:   int32(receiver.psession.SessionID),
		Mtype:       pb.ENMessageType_EN_Message_Request,
		RouteType:   pb.ENRouteType_EN_Route_p2p,
	}
	head := &pb.PBHead{
		Route: route,
		Uid:   receiver.psession.Head_.Uid,
		Cmd:   pb.PBCMsgCmd_ss_request_login,
	}
	msg := &pb.PBCMsg{}
	request := msg.GetSsRequestLogin()
	request.RequestPurpose = EN_Purpose_Get_Data
	request.Uid = receiver.psession.Head_.Uid

	// 获取连接
	socketHandler_ := server.SvrMap[head.Route.Destination]
	sender := message.Message{
		WebSocketHandler: nil,
		SocketHandler:    socketHandler_,
	}
	sender.SendRequestToUser(head, msg)
	return EN_Handler_Succ
}

func (receiver loginProcesser) ProcessResponseMsg() int {
	// 检查用户数据
	ss_response := receiver.psession.ResponseMsg_.GetSsResponseLogin()
	var msg *pb.PBCMsg
	response := *msg.GetCsResponseLogin()
	response.Result = ss_response.Result
	route := &pb.PBRoute{
		Source:      pb.ENPositionType_EN_Position_Connect,
		Destination: pb.ENPositionType_EN_Position_Client,
		SessionId:   int32(receiver.psession.SessionID),
		Mtype:       pb.ENMessageType_EN_Message_Response,
		RouteType:   pb.ENRouteType_EN_Route_p2p,
	}
	head := &pb.PBHead{
		Route: route,
		Uid:   receiver.psession.Head_.Uid,
		Cmd:   pb.PBCMsgCmd_ss_response_login,
	}

	if response.Result == pb.ENMessageError_EN_MESSAGE_ERROR_OK {
		response.User = ss_response.UserData
		websocketHandler := server.ConnectionsMap[ss_response.Uid]
		sender := message.Message{
			WebSocketHandler: websocketHandler,
			SocketHandler:    nil,
		}
		if sender.SendResponseToClient(head, msg) == false {
			return EN_Handler_Done
		}
		// 更新用户位置
		receiver.updateUserPos(ss_response.UserData)
		// 拉取缓存消息
		receiver.getCacheMessage(ss_response.UserData)
		return EN_Handler_Done
	} else {
		// 登录失败
		response.User = nil
		websocketHandler := server.ConnectionsMap[ss_response.Uid]
		sender := message.Message{
			WebSocketHandler: websocketHandler,
			SocketHandler:    nil,
		}
		sender.SendResponseToClient(head, msg)
	}
	return EN_Handler_Done
}

// 更新用户的位置信息[记录登录日志]
func (receiver loginProcesser) updateUserPos(user *pb.PBUser) {
}

// 拉取缓存的消息
func (receiver loginProcesser) getCacheMessage(user *pb.PBUser) {

}

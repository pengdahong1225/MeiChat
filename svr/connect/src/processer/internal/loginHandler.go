package internal

import (
	"connect/src/common/message"
	pb "connect/src/proto"
	"connect/src/server/wsconnect"
)

// 登录请求
type loginProcesser struct {
	processCommon
}

func newLoginProcesser() *loginProcesser {
	return &loginProcesser{}
}

func (receiver loginProcesser) ProcessRequestMsg() int {
	receiver.getUserData(receiver.psession)
	return EN_Handler_Succ
}

func (receiver loginProcesser) ProcessResponseMsg() int {
	// 校验用户数据
	ssResponse := receiver.psession.ResponseMsg_.GetSsResponseGetUserData()

	var msg *pb.PBCMsg
	response := *msg.GetCsResponseLogin()
	response.Result = ssResponse.Result
	route := &pb.PBRoute{
		Source:      pb.ENPositionType_EN_Position_Connect,
		Destination: pb.ENPositionType_EN_Position_Client,
		RouteType:   pb.ENRouteType_EN_Route_p2p,
	}
	head := &pb.PBHead{
		Route:     route,
		Uid:       receiver.psession.Head_.Uid,
		Cmd:       cs_response_login,
		SessionId: int32(receiver.psession.SessionID),
		Mtype:     pb.ENMessageType_EN_Message_Response,
	}

	// 客户端链接
	websocketHandler := wsconnect.ConnectionsMap[ssResponse.Uid]
	if response.Result == pb.ENMessageError_EN_MESSAGE_ERROR_OK {
		response.User = ssResponse.UserData
		if message.SendResponseToClient(websocketHandler, head, msg) == false {
			return EN_Handler_Done
		}
		// 更新用户位置
		receiver.updateUserPos(ssResponse.UserData)
		// 拉取缓存消息
		receiver.getCacheMessage(ssResponse.UserData)
		return EN_Handler_Done
	} else {
		// 登录失败
		response.User = nil
		message.SendResponseToClient(websocketHandler, head, msg)
	}
	return EN_Handler_Done
}

// 更新用户的位置信息[记录登录日志]
func (receiver loginProcesser) updateUserPos(user *pb.PBUser) {
}

// 拉取缓存的消息
func (receiver loginProcesser) getCacheMessage(user *pb.PBUser) {

}

package internal

import (
	"connect/src/common/message"
	pb "connect/src/proto"
)

// 注册请求
type registProcesser struct {
	processCommon
}

func newRegistProcesser() *registProcesser {
	return &registProcesser{}
}

func (receiver registProcesser) ProcessRequestMsg() int {
	receiver.addUserData(receiver.psession)
	return EN_Handler_Succ
}

func (receiver registProcesser) ProcessResponseMsg() int {
	// 校验
	ssResponse := receiver.psession.ResponseMsg_.GetSsResponseAddData()

	var msg *pb.PBCMsg
	response := *msg.GetCsResponseRegist()
	response.Uid = ssResponse.Uid
	response.Result = ssResponse.Result
	route := &pb.PBRoute{
		Source:      pb.ENPositionType_EN_Position_Connect,
		Destination: pb.ENPositionType_EN_Position_Client,
		RouteType:   pb.ENRouteType_EN_Route_p2p,
	}
	head := &pb.PBHead{
		Route:     route,
		Uid:       receiver.psession.Head_.Uid,
		Cmd:       cs_response_regist,
		SessionId: int32(receiver.psession.SessionID),
		Mtype:     pb.ENMessageType_EN_Message_Response,
	}

	message.SendResponseToClient(head, msg)
	return EN_Handler_Done
}

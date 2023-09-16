package internal

import (
	"connect/src/common/message"
	"connect/src/common/session"
	pb "connect/src/proto"
	"connect/src/server/connect"
	"fmt"
)

// ENHandlerResult
const (
	EN_Handler_Done = iota
	EN_Handler_Succ
)

// 状态机 -- 多态处理请求
func ProcessEvent(process_ process) {
	result := EN_Handler_Done
	if process_.GetSession().MessageType_ == pb.ENMessageType_EN_Message_Request {
		process_.GetSession().SessionState_ = session.EN_Session_Idle
		result = process_.ProcessRequestMsg()
	} else if process_.GetSession().MessageType_ == pb.ENMessageType_EN_Message_Response {
		if process_.GetSession().SessionState_ == session.EN_Session_Idle {
			result = process_.ProcessResponseMsg()
		}
	}
	switch result {
	case EN_Handler_Done:
		endProcess(process_.GetSession())
		return
	case EN_Handler_Succ:
		return
	default:
		return
	}
}

// 结束会话
func endProcess(psession *session.Session) {
	session.ManagerInstance.ReleaseSession(psession.SessionID)
	fmt.Printf("session[%d] is release\n", psession.SessionID)
}

// ///////////////////////////////////////////////////////////////////////
// processCommon 公共方法
type processCommon struct {
	psession *session.Session
}

func (receiver processCommon) SetSession(p *session.Session) {
	receiver.psession = p
}

func (receiver processCommon) GetSession() *session.Session {
	return receiver.psession
}

// 获取用户数据
func (receiver processCommon) getUserData(psession *session.Session) {
	route := &pb.PBRoute{
		Source:      pb.ENPositionType_EN_Position_Connect,
		Destination: pb.ENPositionType_EN_Position_User,
		RouteType:   pb.ENRouteType_EN_Route_p2p,
	}
	head := &pb.PBHead{
		Route:     route,
		Uid:       psession.Head_.Uid,
		Cmd:       ss_request_get_user_data,
		SessionId: int32(psession.SessionID),
		Mtype:     pb.ENMessageType_EN_Message_Request,
	}
	msg := &pb.PBCMsg{}
	request := msg.GetSsRequestGetUserData()
	request.Uid = psession.Head_.Uid

	// 获取连接
	socketHandler_ := connect.SvrMap[head.Route.Destination]
	message.SendRequestToUser(socketHandler_.GetConnection(), head, msg)
}

// 更新用户数据
func (receiver processCommon) pushUserData(psession *session.Session) {

}

// 新增用户数据
func (receiver processCommon) addUserData(psession *session.Session) {
	userData := &pb.PBUser{
		Uid:     psession.Head_.Uid,
		Account: psession.RequestMsg_.GetCsRequestRegist().Account,
		Pwd:     psession.RequestMsg_.GetCsRequestRegist().Pwd,
		Gender:  psession.RequestMsg_.GetCsRequestRegist().Gender,
		PicUrl:  "",
	}

	route := &pb.PBRoute{
		Source:      pb.ENPositionType_EN_Position_Connect,
		Destination: pb.ENPositionType_EN_Position_User,
		RouteType:   pb.ENRouteType_EN_Route_p2p,
	}
	head := &pb.PBHead{
		Route:     route,
		Uid:       psession.Head_.Uid,
		Cmd:       ss_request_add_data,
		SessionId: int32(psession.SessionID),
		Mtype:     pb.ENMessageType_EN_Message_Request,
	}
	msg := &pb.PBCMsg{}
	request := msg.GetSsRequestAddData()
	request.Uid = psession.Head_.Uid
	request.UserData = userData

	// 获取连接
	socketHandler_ := connect.SvrMap[head.Route.Destination]
	message.SendRequestToUser(socketHandler_.GetConnection(), head, msg)
}

package internal

import (
	"fmt"
	"user/src/common/db"
	"user/src/common/session"
	pb "user/src/proto"
)

// 用户数据查询服务
type getDataRequest struct {
	session_ *session.Session
}

func NewGetDataRequest() *getDataRequest {
	return &getDataRequest{
		session_: nil,
	}
}

func (receiver getDataRequest) SetSession(psession *session.Session) {
	receiver.session_ = psession
}
func (receiver getDataRequest) GetSession() *session.Session {
	return receiver.session_
}
func (receiver getDataRequest) ProcessRequestMsg() *pb.PBCMsg {
	resMsg := &pb.PBCMsg{}
	response := resMsg.GetSsResponseGetUserData()

	// 拉取用户数据
	user := db.GetData(receiver.session_.Head_.Uid)
	if user == nil {
		response.Result = pb.ENMessageError_EN_MESSAGE_ERROR_INVALID
		return resMsg
	}

	// 验证
	if receiver.session_.Head_.Uid != user.Uid {
		fmt.Printf("request_uid[%d] can't match db_uid[%d]\n", receiver.session_.Head_.Uid, user.Uid)
		response.Result = pb.ENMessageError_EN_MESSAGE_ERROR_INVALID
		return resMsg
	}

	response.Result = pb.ENMessageError_EN_MESSAGE_ERROR_OK
	response.Uid = user.Uid
	response.UserData = user

	return resMsg
}

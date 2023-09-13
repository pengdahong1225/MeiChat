package internal

import (
	"fmt"
	"user/src/common/db"
	"user/src/common/session"
	pb "user/src/proto"
)

// 用户数据查询服务
type loginRequest struct {
	session_ *session.Session
}

func NewLoginRequest() *loginRequest {
	return &loginRequest{
		session_: nil,
	}
}

func (receiver loginRequest) SetSession(psession *session.Session) {
	receiver.session_ = psession
}
func (receiver loginRequest) GetSession() *session.Session {
	return receiver.session_
}
func (receiver loginRequest) ProcessRequestMsg() *pb.PBCMsg {
	resMsg := &pb.PBCMsg{}
	response := resMsg.GetSsResponseLogin()

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

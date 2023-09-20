package server

import (
	"encoding/binary"
	"github.com/duke-git/lancet/v2/slice"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
	"user/src/common"
	"user/src/common/session"
	"user/src/processer"
	pb "user/src/proto"
)

// 新连接处理 - 轮训
func (receiver tcpServer) newConnectionHandle(conn *net.TCPConn) {
	var rbuf = make([]byte, 1024)
	codec := GetCodec()
	for {
		// 接收
		n, err := conn.Read(rbuf)
		if err != nil {
			break
		}
		data := make([]byte, len(rbuf))
		copy(data, rbuf[:n])
		// 数据包解析
		body, errDec := codec.deCode(data)
		if errDec != nil {
			log.Println(errDec)
			break
		}
		// 异步
		bodyTmp := make([]byte, len(body))
		copy(bodyTmp, body)
		receiver.handleAsync(bodyTmp, conn)
	}
	conn.Close()
	receiver.connections = slice.DeleteAt(receiver.connections, slice.IndexOf(receiver.connections, conn))
}

func (receiver tcpServer) handleAsync(body []byte, conn *net.TCPConn) {
	codec := GetCodec()
	err := common.AntsPoolInstance.Submit(func() {
		out := receiver.handleSync(body, conn)
		packet, _ := codec.enCode(out)
		conn.Write(packet)
	})
	if err != nil {
		log.Println(err)
	}
}

func (receiver tcpServer) handleSync(body []byte, conn *net.TCPConn) (out []byte) {
	// 解包
	header, msg, err := deCodeMsg(body)
	if err != nil {
		println(err)
		return nil
	}
	// 处理
	response := receiver.handle(header, msg)
	// 回包
	out, err = receiver.enCodeMsg(header, response)
	if err != nil {
		log.Println(err)
		return nil
	}
	return
}

func (receiver tcpServer) handle(head *pb.PBHead, msg *pb.PBCMsg) *pb.PBCMsg {
	// 分配session
	s := session.ManagerInstance.AllocSession()
	s.Head_ = head
	s.RequestMsg_ = msg
	s.MessageType_ = s.Head_.Mtype
	handler := processer.Instance()
	out := handler.Process(s)
	// 释放session
	session.ManagerInstance.ReleaseSession(s.SessionID)
	return out
}

// ////////////////////////////////////////////////////////////////
func (receiver tcpServer) enCodeMsg(pbHead *pb.PBHead, msg *pb.PBCMsg) ([]byte, error) {
	// route reverse
	pbHead.Route.Source, pbHead.Route.Destination = pbHead.Route.Destination, pbHead.Route.Source

	headBuf := make([]byte, 10)
	if err := proto2Msg(headBuf, pbHead); err != nil {
		return nil, err
	}
	bodyBuf := make([]byte, 10)
	if err := proto2Msg(bodyBuf, msg); err != nil {
		return nil, err
	}

	var ret []byte
	ret = append(ret, headBuf...)
	ret = append(ret, bodyBuf...)

	return ret, nil
}

func deCodeMsg(buf []byte) (*pb.PBHead, *pb.PBCMsg, error) {
	pbHead := &pb.PBHead{}
	if err := msg2Proto(buf, pbHead); err != nil {
		return nil, nil, err
	}
	pbMsg := &pb.PBCMsg{}
	if err := msg2Proto(buf, pbMsg); err != nil {
		return nil, nil, err
	}
	return pbHead, pbMsg, nil
}
func proto2Msg(buf []byte, message proto.Message) error {
	msgBuf, err := proto.Marshal(message)
	msgLen := len(msgBuf)
	lenBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(lenBuf, uint32(msgLen))

	buf = append(buf, lenBuf...)
	buf = append(buf, msgBuf...)
	return err
}
func msg2Proto(buf []byte, message proto.Message) error {
	// 先读取这部分的长度
	lenBuf := buf[:4]
	msgLen := binary.BigEndian.Uint32(lenBuf)

	msgBuf := buf[4:int(msgLen)]
	return proto.Unmarshal(msgBuf, message)
}

// ////////////////////////////////////////////////////////////////

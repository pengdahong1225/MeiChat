package codec

import (
	pb "connect/src/proto"
	"encoding/binary"
	"google.golang.org/protobuf/proto"
)

func (receiver Codec) DeCodeMsg(buf []byte) (*pb.PBHead, *pb.PBCMsg, error) {
	pbHead := &pb.PBHead{}
	if err := receiver.msg2Proto(buf, pbHead); err != nil {
		return nil, nil, err
	}
	pbMsg := &pb.PBCMsg{}
	if err := receiver.msg2Proto(buf, pbMsg); err != nil {
		return nil, nil, err
	}
	return pbHead, pbMsg, nil
}
func (receiver Codec) msg2Proto(buf []byte, message proto.Message) error {
	// 先读取这部分的长度
	lenBuf := buf[:4]
	msgLen := binary.BigEndian.Uint32(lenBuf)

	msgBuf := buf[4:int(msgLen)]
	return proto.Unmarshal(msgBuf, message)
}

func (receiver Codec) EnCodeMsg(head *pb.PBHead, msg *pb.PBCMsg) ([]byte, error) {
	headBuf := make([]byte, 10)
	if err := proto2Msg(headBuf, head); err != nil {
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
func proto2Msg(buf []byte, message proto.Message) error {
	msgBuf, err := proto.Marshal(message)
	msgLen := len(msgBuf)
	lenBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(lenBuf, uint32(msgLen))

	buf = append(buf, lenBuf...)
	buf = append(buf, msgBuf...)
	return err
}

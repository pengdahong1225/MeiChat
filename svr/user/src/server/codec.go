package server

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"strconv"
)

const (
	packetSize    = 4    // 包长
	magicCode     = "XX" // 魔术码
	magicCodeSize = 2    // 魔术码长度
)

// Protocol format:
//
// * 0                       4           6
// * +-----------------------+-----------+
// * |   packet len          |magic code |
// * +-----------+-----------+-----------+
// * |                                   |
// * +                                   +
// * |           body bytes              |
// * +                                   +
// * |            ... ...                |
// * +-----------------------------------+.

type Codec struct {
}

func GetCodec() *Codec {
	return &Codec{}
}

// 数据包
func (receiver Codec) enCode(buf []byte) ([]byte, error) {
	data := make([]byte, packetSize+magicCodeSize+len(buf))
	// head
	binary.BigEndian.PutUint32(data[0:packetSize], uint32(len(buf)))
	copy(data[packetSize:packetSize+magicCodeSize], magicCode)
	// body
	copy(data[packetSize+magicCodeSize:len(buf)], buf)
	return data, nil
}
func (receiver Codec) deCode(buf []byte) ([]byte, error) {
	head := buf[:packetSize+magicCodeSize]
	if len(head) != packetSize+magicCodeSize {
		return nil, errors.New("invalid head")
	}

	if !checkHead(head) {
		return nil, errors.New("invalid magic")
	}

	reader := bytes.NewReader(buf[packetSize+magicCodeSize:])
	body, err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.New("invalid body")
	}

	if !checkBody(head, len(body)) {
		return nil, errors.New("head can't match body")
	}
	return body, nil
}

// 验证
func checkHead(head []byte) bool {
	magic := head[packetSize:]
	if !bytes.Equal([]byte(magicCode), magic) {
		return false
	}
	return true
}

// 验证头部消息
func checkBody(head []byte, bodyLength int) bool {
	header := head[:packetSize]
	lens, _ := strconv.ParseInt(string(header), 10, 32)
	return lens == int64(bodyLength)
}

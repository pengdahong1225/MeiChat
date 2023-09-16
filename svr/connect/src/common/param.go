package common

import (
	pb "connect/src/proto"
	"github.com/gorilla/websocket"
	"net"
)

var ConnectionsMap map[int64]*websocket.Conn // (uid,客户端连接)
var SvrMap map[pb.ENPositionType]*net.Conn

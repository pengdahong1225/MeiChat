package initializer

import (
	"connect/src/common"
	pb "connect/src/proto"
	"fmt"
	"github.com/spf13/viper"
	"net"
	"strconv"
)

// webServer.conf
type serverInfo struct {
	User_ user `mapstructure:"user"`
	Chat_ chat `mapstructure:"chatServer"`
}
type user struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
type chat struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

var serverInfoInstance serverInfo

func init() {
	// 初始化conf
	serverInfoInstance = initServerConf()
	// 初始化链接handler
	initServer()
}

func initServerConf() serverInfo {
	viperConfig := viper.New()
	viperConfig.SetConfigName("webServer.conf")
	viperConfig.AddConfigPath("conf") // 容器入口点的相对路径
	viperConfig.SetConfigType("ini")

	err := viperConfig.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 反射
	config := serverInfo{}
	if err := viperConfig.Unmarshal(&config); err != nil {
		panic(err)
	}
	return config
}

func initServer() {
	common.SvrMap = make(map[pb.ENPositionType]*net.Conn)
	if common.SvrMap[pb.ENPositionType_EN_Position_User] = newTcpHandler(serverInfoInstance.User_.Host,
		serverInfoInstance.User_.Port); common.SvrMap[pb.ENPositionType_EN_Position_User] == nil {
		fmt.Println("tcpconnect to user error")
		return
	}
	if common.SvrMap[pb.ENPositionType_EN_Position_ChatServer] = newTcpHandler(serverInfoInstance.Chat_.Host,
		serverInfoInstance.Chat_.Port); common.SvrMap[pb.ENPositionType_EN_Position_ChatServer] == nil {
		fmt.Println("tcpconnect to chatServer error")
		return
	}
}

func newTcpHandler(ip string, port int) *net.Conn {
	conn, err := net.Dial("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &conn
}

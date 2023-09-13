package initializer

import (
	"github.com/spf13/viper"
)

// server.conf
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

var ServerInfoInstance serverInfo

func init() {
	ServerInfoInstance = initServer()
}

func initServer() serverInfo {
	viperConfig := viper.New()
	viperConfig.SetConfigName("server.conf")
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

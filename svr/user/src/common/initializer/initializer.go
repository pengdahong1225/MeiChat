package initializer

import (
	"github.com/spf13/viper"
)

// db.conf
type dbInfo struct {
	Sql_   sqlInfo   `mapstructure:"mysql"`
	NoSql_ redisInfo `mapstructure:"redis"`
}
type sqlInfo struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	User string `mapstructure:"user"`
	Pwd  string `mapstructure:"pwd"`
	DB   string `mapstructure:"db_name"`
}
type redisInfo struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

var DBInfoInstance dbInfo

func init() {
	DBInfoInstance = initDB()
}

func initDB() dbInfo {
	viperConfig := viper.New()
	viperConfig.SetConfigName("db.conf")
	viperConfig.AddConfigPath("conf") // 容器入口点的相对路径
	viperConfig.SetConfigType("ini")

	err := viperConfig.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 反射
	config := dbInfo{}
	if err := viperConfig.Unmarshal(&config); err != nil {
		panic(err)
	}
	return config
}

package initialize

import (
	"fmt"
	"grpctest/global"

	"github.com/spf13/viper"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	//从配置文件中读取出对应的配置
	debug := GetEnvInfo("MXSHOP_DEBUG")
	configFilePrefix:="config"
	configFileName:=fmt.Sprintf("../grpctest/%s-pro.yaml",configFilePrefix)
	if debug{
		configFileName=fmt.Sprintf("../grpctest/%s-debug.yaml",configFilePrefix)
	}
	v:=viper.New()
	v.SetConfigFile(configFileName)
	if err:=v.ReadInConfig();err!=nil{
		panic(err)
	}
	if err:=v.Unmarshal(&global.ServerConfig);err!=nil{
		panic(err)
	}
}


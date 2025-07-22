package main

import (
	"flag"
	"fmt"
	_ "grpctest/global"
	"grpctest/handler"
	"grpctest/initialize"
	"grpctest/model/proto"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	IP:=flag.String("ip","0.0.0.0","ip地址")
	Port:=flag.Int("port",50051,"端口号")

	//初始化	
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()


	flag.Parse()
	zap.S().Info("ip:",*IP)
	fmt.Println("port:",*Port)
	server := grpc.NewServer()
	proto.RegisterUserServer(server,&handler.UserServer{})
	// lis,err:=net.Listen("tcp", "0.0.0.0:8088")
	lis,err:=net.Listen("tcp",fmt.Sprintf("%s:%d",*IP,*Port))
	if err!=nil{
		panic("failed to listen:"+err.Error())
	}
	err=server.Serve(lis)
	if err!=nil{
		panic("failed to start grpc:"+err.Error()                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          )
	}
}
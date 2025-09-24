package main

import (
	"flag"
	"fmt"
	"grpctest/global"
	_ "grpctest/global"
	"grpctest/handler"
	"grpctest/initialize"
	"grpctest/model/proto"
	"grpctest/utils"
	"net"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	IP:=flag.String("ip","0.0.0.0","ip地址")
	Port:=flag.Int("port",0,"端口号")

	//初始化	
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	zap.S().Info(global.ServerConfig)


	flag.Parse()
	zap.S().Info("ip:",*IP)
	fmt.Println("port:",*Port)
	//如果端口为0，随机获取一个端口
	if *Port==0{
		port,err:=utils.GetFreePort()
		if err!=nil{
			panic(err)
		}
		*Port=port
	}
	
	server := grpc.NewServer()
	proto.RegisterUserServer(server,&handler.UserServer{})
	// lis,err:=net.Listen("tcp", "0.0.0.0:8088")
	lis,err:=net.Listen("tcp",fmt.Sprintf("%s:%d",*IP,*Port))
	if err!=nil{
		panic("failed to listen:"+err.Error())
	}
	//注册健康检查
	grpc_health_v1.RegisterHealthServer(server,health.NewServer())
	
	//服务注册
	cfg:=api.DefaultConfig()//获取 Consul 客户端的默认配置
	cfg.Address=fmt.Sprintf("%s:%d",global.ServerConfig.ConsulInfo.Host,global.ServerConfig.ConsulInfo.Port)//这是 Consul 服务器的地址。


	client,err:=api.NewClient(cfg)//创建一个 Consul 客户端实例
	if err!=nil{
		panic(err)
	}
	check:=&api.AgentServiceCheck{
		//HTTP: "http://127.0.0.1:8021/health",
		GRPC: fmt.Sprintf("192.168.11.120:%d",*Port),//Consul 会定期请求这个 URL 来检查服务是否健康
		Interval: "5s",//每 5s 检查一次。
    	Timeout: "2s",// 检查响应超时时间为 2s。
    	DeregisterCriticalServiceAfter: "10s",//如果连续失败达 10s，则从注册列表中移除服务。
	}
	//生成注册对象
	registration:=new(api.AgentServiceRegistration)
	registration.Name=global.ServerConfig.Name//服务名称（如 "user-web"）。
	registration.ID=global.ServerConfig.Name//服务实例唯一标识（如 "user-web"，如果部署多个副本，每个副本需要不同 ID）。
	registration.Port=*Port//服务监听的端口号（如 8021）
	registration.Tags=[]string{"mxshop","boddy","imooc","user","srv"}//服务标签，用于分类或标记服务（如 ["mxshop", "boddy"]）。
	registration.Address="192.168.11.120"//服务所在主机的 IP 地址,会被web层拿到做服务发现，consul自己来做检查的
	registration.Check=check
	if err := client.Agent().ServiceRegister(registration); err != nil {
    zap.S().Fatalf("register service to consul failed: %v", err)
}
	err=server.Serve(lis)
	if err!=nil{
		panic("failed to start grpc:"+err.Error()                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          )
	}
}
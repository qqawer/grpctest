package main

import (
	"context"
	"fmt"
	"grpctest/model/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var userClient proto.UserClient
var conn *grpc.ClientConn


func Init() {
	var err error
	conn, err = grpc.NewClient(
		"127.0.0.1:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil{
		panic(err)
	}
	// defer conn.Close()

	userClient = proto.NewUserClient(conn)
	
}

func TestCreateUser(){
	for i:=0;i<10;i++{
		rsp,err:=userClient.CreateUser(context.Background(),&proto.CreateUserInfo{
			NickName:fmt.Sprintf("bobby%d", i),
			PassWord: "admin123",
			Mobile: fmt.Sprintf("1878222222%d",i),
		})

		if err!=nil{
			panic(err)
		}
		fmt.Println(rsp.Id)
	}
	
}
func main() {
	Init()
	TestCreateUser()
	conn.Close()
}
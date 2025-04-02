package main

import (
	"grpctest/global"
	model "grpctest/model/user"
)

func main() {
	
	//设置全局的logger,这个logger在我们执行每个sql语句的时候会打印每一行sql
	//sql才是最重要的，本着这个的原则我尽量的给大家看到每个api背后的sql语句是什么

	//定义一个表结构，将表结构直接生成对应的表 - migrations
	//迁移schema
	
  // Migrate the schema
  
  global.DB.AutoMigrate(&model.User{})
}
package main

import (
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"wendaxitong/user/internal/handler"
	"wendaxitong/user/internal/repository"
	"wendaxitong/user/internal/service"
)

func main() {
	err := repository.ConnectMysqlDatabase() // 连接数据库
	if err != nil {
		fmt.Println("连接数据库失败:", err.Error())
	}
	err = StartGRPCServiceServer() // 启动服务端grpc服务
	if err != nil {
		fmt.Println("启动服务端grpc服务失败:", err.Error())
	}
}

func StartGRPCServiceServer() error {
	// 启动服务端grpc服务
	rpcServer := grpc.NewServer()
	service.RegisterUserServiceServer(rpcServer, handler.NewUserService())
	// 启动监听
	listener, err := net.Listen("tcp", ":8070")
	if err != nil {
		return errors.New("启动监听错误：" + err.Error())
	}
	// 启动服务
	err = rpcServer.Serve(listener)
	if err != nil {
		return errors.New("启动服务错误：" + err.Error())
	}
	return nil
}

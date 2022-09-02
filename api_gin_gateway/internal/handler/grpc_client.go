package handler

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"wendaxitong/api_gin_gateway/internal/service"
)

var GrpcUerServiceClient service.UserServiceClient

func StartGRPCClientServer() {
	conn, err := grpc.Dial(":8070", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("服务端出错，连接不上：", err)
	}
	GrpcUerServiceClient = service.NewUserServiceClient(conn)
}

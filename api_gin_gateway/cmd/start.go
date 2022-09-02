package main

import (
	"wendaxitong/api_gin_gateway/internal/handler"
	"wendaxitong/api_gin_gateway/pkg/util"
	"wendaxitong/api_gin_gateway/routers"
)

func main() {
	util.ConnectRedis()             // 连接redis
	handler.StartGRPCClientServer() // 启动客户端grpc
	r := routers.NewRouter()        // 初始化路由
	r.Run(":8080")

}

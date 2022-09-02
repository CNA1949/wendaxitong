package main

/*
	该文件用于测试
*/
import (
	"wendaxitong/api_gin_gateway/pkg/util"
)

func main() {
	util.ConnectRedis() // 连接redis
	//util.SetKeyValue("name", "zhong", 300)
	util.DeleteKeyValue("name")
}

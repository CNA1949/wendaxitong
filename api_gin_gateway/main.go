package main

/*
	该文件用于测试
*/
import (
	"fmt"
	"wendaxitong/api_gin_gateway/pkg/util"
)

func main() {
	util.ConnectRedis() // 连接redis
	util.SetKeyValue("name", "tan", 300)
	v, _ := util.GetValueByKey("name")
	fmt.Println(v)
}

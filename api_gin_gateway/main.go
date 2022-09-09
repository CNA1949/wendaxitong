package main

import (
	"fmt"
	"wendaxitong/api_gin_gateway/pkg/util"
)

func main() {
	deleteAllKeyValue()
}
func deleteAllKeyValue() {
	util.ConnectRedis()
	err := util.DeleteAllKeyValue()
	if err != nil {
		fmt.Println(err.Error())
	}
}

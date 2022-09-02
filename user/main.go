package main

import (
	"fmt"
	"wendaxitong/user/config"
)

func main() {
	var config config.Configuration
	config.GetMysqlConfig()
	fmt.Println(config.MysqlConfig)
}

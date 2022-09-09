package main

/*
	该文件用于测试
*/
import (
	"fmt"
	"wendaxitong/user/internal/repository"
)

func main() {
	repository.ConnectMysqlDatabase()
	err := repository.UpdateValueByName("user_name", ",cna2", &repository.UserInfo{}, "num_fans", 4)
	fmt.Println(err)
}

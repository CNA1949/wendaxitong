package util

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"wendaxitong/api_gin_gateway/config"
)

var ConnRedis redis.Conn

func ConnectRedis() {
	// 连接redis	,必须手动先启动redis
	var config config.Configuration
	config.GetMysqlConfig()
	redisConfig := config.RedisConfig
	// 连接redis
	var err error
	ConnRedis, err = redis.Dial("tcp", redisConfig.Host+":"+redisConfig.Port)
	if err != nil {
		fmt.Println("Connect to redis error:", err)
		return
	} else {
		fmt.Println("Connect to redis successfully")
	}

	// 密码鉴权
	_, err = ConnRedis.Do("AUTH", redisConfig.Password)
	if err != nil {
		fmt.Println("auth failed:", err)
	} else {
		fmt.Println("auth ok")
	}

}

// SetKeyValue 设置键值对
func SetKeyValue(key string, value string, expireTime int) error {
	_, err := ConnRedis.Do("set", key, value)
	if err != nil {
		fmt.Println("err:", err.Error())
		return err
	}
	_, err = ConnRedis.Do("expire", key, expireTime) // 设置键值对过期时间，expireTime单位为秒
	if err != nil {
		fmt.Println("err:", err.Error())
		return err
	}
	return nil
}

// GetValueByKey 获取key的value
func GetValueByKey(key string) (string, error) {
	value, err := redis.String(ConnRedis.Do("get", key))
	if err != nil {
		fmt.Println("err:", err.Error())
		return "", err
	}
	return value, nil
}

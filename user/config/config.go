package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Configuration struct {
	MysqlConfig MysqlConfiguration `yaml:"mysql" json:"mysql_config"`
	RedisConfig RedisConfiguration `yaml:"redis" json:"redis_config"`
}

type MysqlConfiguration struct {
	DriverName string `yaml:"driverName" json:"driver_name"`
	UserName   string `yaml:"userName" json:"user_name"`
	Password   string `yaml:"password" json:"password"`
	Protocol   string `yaml:"protocol" json:"protocol"`
	Host       string `yaml:"host" json:"host"`
	Port       string `yaml:"port" json:"port"`
	Database   string `yaml:"database" json:"database"`
	Charset    string `yaml:"charset" json:"charset"`
	ParseTime  string `yaml:"parseTime" json:"parse_time"`
	Loc        string `yaml:"loc" json:"loc"`
}

type RedisConfiguration struct {
	Host     string `yaml:"host" json:"host"`
	Port     string `yaml:"port" json:"port"`
	Password string `yaml:"password" json:"password"`
}

func (config *Configuration) GetConfig() {
	// 获取配置信息
	workDir, _ := os.Getwd()
	yamlFile, err := ioutil.ReadFile(workDir + "\\user\\config\\config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	// 将配置文件读取到结构体中
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		fmt.Println(err.Error())
	}
}

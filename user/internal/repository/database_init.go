package repository

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"wendaxitong/user/config"
)

var DB *gorm.DB

const (
	USERNAME = "root"                                     //数据库用户名
	PASSWORD = "123456"                                   // 数据库登录密码
	PROTOCOL = "tcp"                                      //协议
	ADDRESS  = "localhost:3306"                           //地址
	DBNAME   = "wendaxitong"                              //数据库名
	PARAMS   = "charset=utf8mb4&parseTime=True&loc=Local" //其它参数
)

func ConnectMysqlDatabase() error {
	var config config.Configuration
	config.GetConfig()
	mysqlConfig := config.MysqlConfig // 获取mysql数据库配置信息
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		mysqlConfig.UserName, mysqlConfig.Password, mysqlConfig.Protocol,
		mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database,
		mysqlConfig.Charset, mysqlConfig.ParseTime, mysqlConfig.Loc)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := sql.Open(mysqlConfig.DriverName, dsn)
	if err != nil {
		return err
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
	return nil
}

// UpdateValueById 根据ID，更新某个字段的值
func UpdateValueById(idName string, id interface{}, model interface{}, attribute string, newValue interface{}) error {
	query := idName + " = ?"
	tx := DB.Begin() //开启事务
	err := DB.Model(&model).Where(query, id).Update(attribute, newValue).Error
	if err != nil {
		tx.Rollback() // 遇到错误时回滚事务
		return err
	}
	tx.Commit() // 提交事务
	return nil
}

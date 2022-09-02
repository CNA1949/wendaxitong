package repository

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
	"wendaxitong/user/internal/service"
	"wendaxitong/user/pkg/codeMsg"
)

type UserInfo struct {
	UserId        uint      `gorm:"primary_key;auto_increment" json:"user_id"`
	UserName      string    `gorm:"not null;unique" json:"user_name"`
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
	NumFans       uint      `json:"num_fans"`
	NumIdols      uint      `json:"num_idols"`
	Password      string    `gorm:"not null" json:"password"`
	CreatedAt     time.Time `json:"created_at"`      // 用户注册时间
	UpdatedAt     time.Time `json:"updated_at"`      // 用户信息更新时间
	LastLoginTime time.Time `json:"last_login_time"` // 用户最近登录时间
}

const (
	PassWordCost = 12 // 密码加密难度
)

// SetPassword 加密密码
func (user *UserInfo) SetPassword(password string) (bool, codeMsg.CodeMessage) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return false, codeMsg.CodeMessage{StatusCode: codeMsg.ErrorEncryptionFailed, StatusMessage: codeMsg.GetErrorMsg(codeMsg.ErrorEncryptionFailed)} // 加密失败
	}
	user.Password = string(bytes)
	return true, codeMsg.CodeMessage{}
}

// CheckPassword 检验密码
func (user *UserInfo) CheckPassword(password string) (bool, codeMsg.CodeMessage) {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println("CheckPassword:", err)
		return false, codeMsg.CodeMessage{StatusCode: codeMsg.FALSE, StatusMessage: "密码错误"}
	} else {
		return true, codeMsg.CodeMessage{StatusCode: codeMsg.TRUE, StatusMessage: "密码正确"}
	}
}

// CheckUserExist 检查用户是否存在
func (user *UserInfo) CheckUserExist(request *service.UserRequest) bool {
	//err := DB.Where("user_name = ?", request.UserName).First(&user).Error
	//if err == gorm.ErrRecordNotFound {
	//	return false
	//}
	var count int64
	var userInfo UserInfo
	DB.Where("user_name = ?", request.UserName).First(&userInfo).Count(&count)
	fmt.Println("count:", count)
	if count != 0 {
		return true
	}
	return false
}

// RegisterUserInfo 用户注册
func (user *UserInfo) RegisterUserInfo(request *service.UserRequest) codeMsg.CodeMessage {
	isExist := user.CheckUserExist(request)
	if isExist {
		return codeMsg.CodeMessage{StatusCode: codeMsg.ErrorUserExist, StatusMessage: codeMsg.GetErrorMsg(codeMsg.ErrorUserExist)} // 该用户已存在
	}
	{
		user.UserName = request.UserName
		user.Phone = request.Phone
		user.Email = request.Email
		user.NumFans = 0
		user.NumIdols = 0
	}

	b, msg := user.SetPassword(request.Password) // 密码加密
	if !b {
		return msg // 加密失败
	}

	// 注册用户信息
	tx := DB.Begin() //开启事务
	err := tx.Create(&user).Error
	if err != nil {
		tx.Rollback()                                                                                                                        // 遇到错误时回滚事务
		return codeMsg.CodeMessage{StatusCode: codeMsg.ErrorRegisterFailed, StatusMessage: codeMsg.GetErrorMsg(codeMsg.ErrorRegisterFailed)} // 注册失败
	}
	tx.Commit() // 提交事务
	return codeMsg.CodeMessage{StatusCode: codeMsg.SUCCESS, StatusMessage: "注册成功"}
}

// UserLogin 用户登录
func (user *UserInfo) UserLogin(request *service.UserRequest) codeMsg.CodeMessage {
	var count int64
	isExist := user.CheckUserExist(request)
	if !isExist {
		return codeMsg.CodeMessage{StatusCode: codeMsg.ErrorUserNotExist, StatusMessage: codeMsg.GetErrorMsg(codeMsg.ErrorUserNotExist)} // 该用户已存在
	}
	DB.Where("user_name = ?", request.UserName).First(&user).Count(&count)

	if count != 0 {
		b, _ := user.CheckPassword(request.Password)
		if b {
			tx := DB.Begin()
			err := tx.Model(&UserInfo{}).Where("user_name = ?", request.UserName).Update("last_login_time", time.Now()).Error
			if err != nil {
				tx.Rollback()                                                                                // 遇到错误时回滚事务
				return codeMsg.CodeMessage{StatusCode: codeMsg.ErrorLoginFailed, StatusMessage: err.Error()} // 登录失败，操作数据库错误
			}
			tx.Commit() // 提交事务
			return codeMsg.CodeMessage{StatusCode: codeMsg.SUCCESS, StatusMessage: "登录成功"}
		} else {
			return codeMsg.CodeMessage{StatusCode: codeMsg.ErrorLoginFailed, StatusMessage: "登录失败，密码错误"}
		}
	} else {
		return codeMsg.CodeMessage{StatusCode: codeMsg.ErrorLoginFailed, StatusMessage: "登录失败，账号不存在"}
	}

}

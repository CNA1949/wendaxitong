package repository

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
	"wendaxitong/user/internal/service"
	"wendaxitong/user/pkg/codeMsg"
)

type UserInfo struct {
	UserId        uint64    `gorm:"primary_key;auto_increment" json:"user_id"`
	UserName      string    `gorm:"not null;unique" json:"user_name"`
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
	NumFans       uint64    `json:"num_fans"`
	NumIdols      uint64    `json:"num_idols"`
	Password      string    `gorm:"not null" json:"password"`
	CreatedAt     time.Time `json:"created_at"`      // 用户注册时间
	UpdatedAt     time.Time `json:"updated_at"`      // 用户信息更新时间
	LastLoginTime time.Time `json:"last_login_time"` // 用户最近登录时间
	FansNames     string    `json:"fans_names"`
	IdolsNames    string    `json:"idols_names"`
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

/**
***
*** 基础功能****************************************************************************
***
 */

// CheckUserExist 检查用户是否存在
func (user *UserInfo) CheckUserExist(userName string) bool {
	var u UserInfo
	//err := DB.Where("user_name = ?", request.UserName).First(&user).Error
	//if err == gorm.ErrRecordNotFound {
	//	return false
	//}
	var count int64
	DB.Where("user_name = ?", userName).First(&u).Count(&count)
	if count != 0 {
		return true
	}
	return false
}

// RegisterUserInfo 用户注册
func (user *UserInfo) RegisterUserInfo(request *service.UserRequest) codeMsg.CodeMessage {
	isExist := user.CheckUserExist(request.UserInfo.UserName)
	if isExist {
		return codeMsg.CodeMessage{StatusCode: codeMsg.ErrorUserExist, StatusMessage: codeMsg.GetErrorMsg(codeMsg.ErrorUserExist)} // 该用户已存在
	}
	{
		user.UserName = request.UserInfo.UserName
		user.Phone = request.UserInfo.Phone
		user.Email = request.UserInfo.Email
		user.NumFans = 0
		user.NumIdols = 0
	}

	b, msg := user.SetPassword(request.UserInfo.Password) // 密码加密
	if !b {
		return msg // 加密失败
	}

	// 注册用户信息
	tx := DB.Begin() //开启事务
	err := tx.Create(user).Error
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
	isExist := user.CheckUserExist(request.UserInfo.UserName)
	if !isExist {
		return codeMsg.CodeMessage{StatusCode: codeMsg.ErrorUserNotExist, StatusMessage: codeMsg.GetErrorMsg(codeMsg.ErrorUserNotExist)} // 该用户不存在
	}
	DB.Where("user_name = ?", request.UserInfo.UserName).First(user).Count(&count)

	if count != 0 {
		b, _ := user.CheckPassword(request.UserInfo.Password)
		if b {
			tx := DB.Begin()
			err := tx.Model(&UserInfo{}).Where("user_name = ?", request.UserInfo.UserName).Update("last_login_time", time.Now()).Error
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

// DeleteUser 注销用户
func (user *UserInfo) DeleteUser(request *service.UserRequest) codeMsg.CodeMessage {
	isExist := user.CheckUserExist(request.UserInfo.UserName)
	if !isExist {
		return codeMsg.CodeMessage{StatusCode: codeMsg.ErrorUserNotExist, StatusMessage: codeMsg.GetErrorMsg(codeMsg.ErrorUserNotExist)} // 该用不存在
	}

	err := DB.Where("user_name = ?", request.UserInfo.UserName).First(user).Error
	tx := DB.Begin() //开启事务
	err = tx.Unscoped().Delete(user).Error
	if err != nil {
		tx.Rollback() // 遇到错误时回滚事务
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "删除失败"}
	}
	tx.Commit() // 提交事务
	return codeMsg.CodeMessage{StatusCode: codeMsg.SUCCESS, StatusMessage: "删除成功"}
}

// ModifyUserInfo 修改用户个人信息
func (user *UserInfo) ModifyUserInfo(request *service.UserRequest) codeMsg.CodeMessage {
	var err error

	if request.UserInfo.UserName != "" {
		isExist := user.CheckUserExist(request.UserInfo.UserName)
		if isExist {
			return codeMsg.CodeMessage{StatusCode: codeMsg.ErrorUserExist, StatusMessage: codeMsg.GetErrorMsg(codeMsg.ErrorUserExist)} // 该用已存在
		}
		err = UpdateValueById("user_id", request.UserInfo.UserId, &UserInfo{}, "user_name", request.UserInfo.UserName)
		if err != nil {
			return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "update user_name:" + err.Error()}
		}
	}

	if request.UserInfo.Password != "" {
		user.SetPassword(request.UserInfo.Password)
		err = UpdateValueById("user_id", request.UserInfo.UserId, &UserInfo{}, "password", user.Password)
		if err != nil {
			return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "update password:" + err.Error()}
		}
	}

	if request.UserInfo.Phone != "" {
		err = UpdateValueById("user_id", request.UserInfo.UserId, &UserInfo{}, "phone", request.UserInfo.Phone)
		if err != nil {
			return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "update phone:" + err.Error()}
		}
	}

	if request.UserInfo.Email != "" {
		err = UpdateValueById("user_id", request.UserInfo.UserId, &UserInfo{}, "email", request.UserInfo.Email)
		if err != nil {
			return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "update email:" + err.Error()}
		}
	}

	msg := user.GetUserInfoByUserId(request)
	if msg.StatusCode != codeMsg.SUCCESS {
		return msg
	}

	return codeMsg.CodeMessage{StatusCode: codeMsg.SUCCESS, StatusMessage: "修改成功"}
}

// GetUserInfoByUserName 根据用户名获取用户信息
func (user *UserInfo) GetUserInfoByUserName(request *service.UserRequest) codeMsg.CodeMessage {
	isExist := user.CheckUserExist(request.UserInfo.UserName)
	if !isExist {
		return codeMsg.CodeMessage{StatusCode: codeMsg.ErrorUserNotExist, StatusMessage: codeMsg.GetErrorMsg(codeMsg.ErrorUserNotExist)} // 该用不存在
	}
	err := DB.Where("user_name = ?", request.UserInfo.UserName).First(user).Error
	if err != nil {
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: err.Error()}
	}
	return codeMsg.CodeMessage{StatusCode: codeMsg.SUCCESS, StatusMessage: "获取成功"}
}

// GetUserInfoByUserId 根据用户ID获取用户信息
func (user *UserInfo) GetUserInfoByUserId(request *service.UserRequest) codeMsg.CodeMessage {
	var count int64
	err := DB.Where("user_id = ?", request.UserInfo.UserId).First(user).Count(&count).Error
	if count == 0 {
		return codeMsg.CodeMessage{StatusCode: codeMsg.ErrorUserNotExist, StatusMessage: codeMsg.GetErrorMsg(codeMsg.ErrorUserNotExist)} // 该用不存在
	}

	if err != nil {
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: err.Error()}
	}
	return codeMsg.CodeMessage{StatusCode: codeMsg.SUCCESS, StatusMessage: "获取成功"}
}

/**
***
*** 用户关注与取消关注****************************************************************************
***
 */

// GetIdolsNames 获取所有关注用户的用户名(map数据类型)
func (user *UserInfo) GetIdolsNames(userName string) (map[string]interface{}, error) {
	var u UserInfo
	isExist := user.CheckUserExist(userName)
	if !isExist {
		return nil, errors.New("该用户不存在")
	}
	err := DB.Where("user_name = ?", userName).First(&u).Error
	if err != nil {
		return nil, err
	}
	u.IdolsNames = strings.Trim(u.IdolsNames, ",")
	names := strings.Split(u.IdolsNames, ",")
	var namesMap = make(map[string]interface{}, 10)
	for i := 0; i < len(names); i++ {
		namesMap[names[i]] = ""
	}
	return namesMap, nil
}

// GetFansNames 获取所有粉丝的用户名(map数据类型)
func (user *UserInfo) GetFansNames(userName string) (map[string]interface{}, error) {
	var u UserInfo
	isExist := user.CheckUserExist(userName)
	if !isExist {
		return nil, errors.New("该用户不存在")
	}
	err := DB.Where("user_name = ?", userName).First(&u).Error
	if err != nil {
		return nil, err
	}
	u.FansNames = strings.Trim(u.FansNames, ",")
	names := strings.Split(u.FansNames, ",")
	var namesMap = make(map[string]interface{}, 10)
	for i := 0; i < len(names); i++ {
		namesMap[names[i]] = ""
	}
	return namesMap, nil
}

// addIdol 增加idol
func (user *UserInfo) addIdol(userName string, idolName string) error {
	var u UserInfo
	err := DB.Where("user_name = ?", userName).First(&u).Error
	if err != nil {
		return err
	}
	newIdolsInfo := u.IdolsNames + "," + idolName
	newIdolsInfo = strings.Trim(newIdolsInfo, ",")
	err = UpdateValueByName("user_name", userName, UserInfo{}, "idols_names", newIdolsInfo)
	if err != nil {
		return err
	}
	err = UpdateValueByName("user_name", userName, UserInfo{}, "num_idols", u.NumIdols+1)
	if err != nil {
		return err
	}
	return nil
}

// addFan 增加fans
func (user *UserInfo) addFan(userName string, fanName string) error {
	userName = strings.Trim(userName, ",")
	fmt.Println(userName)
	var u UserInfo
	err := DB.Where("user_name = ?", userName).First(&u).Error
	if err != nil {
		return err
	}

	newFansInfo := u.FansNames + "," + fanName
	newFansInfo = strings.Trim(newFansInfo, ",")
	err = UpdateValueByName("user_name", userName, UserInfo{}, "fans_names", newFansInfo)
	if err != nil {
		return err
	}
	err = UpdateValueByName("user_name", userName, UserInfo{}, "num_fans", u.NumFans+1)
	if err != nil {
		return err
	}
	return nil
}

// deleteIdol 减除idol
func (user *UserInfo) deleteIdol(userName string, idolName string) error {
	var u UserInfo
	err := DB.Where("user_name = ?", userName).First(&u).Error
	if err != nil {
		return err
	}

	m, err := user.GetIdolsNames(userName)
	if err != nil {
		return err
	}
	delete(m, idolName)
	var newIdolInfo string
	for k, _ := range m {
		newIdolInfo = newIdolInfo + "," + k
	}
	newIdolInfo = strings.Trim(newIdolInfo, ",")
	err = UpdateValueByName("user_name", userName, &UserInfo{}, "idols_names", newIdolInfo)
	if err != nil {
		return err
	}
	err = UpdateValueByName("user_name", userName, &UserInfo{}, "num_idols", u.NumIdols-1)
	if err != nil {
		return err
	}
	return nil
}

// deleteFans 减除fan
func (user *UserInfo) deleteFan(userName string, fanName string) error {
	var u UserInfo
	userName = strings.Trim(userName, ",")
	err := DB.Where("user_name = ?", userName).First(&u).Error
	if err != nil {
		return err
	}

	m, err := user.GetFansNames(userName)
	if err != nil {
		return err
	}
	delete(m, fanName)
	var newFanInfo string
	for k, _ := range m {
		newFanInfo = newFanInfo + "," + k
	}
	newFanInfo = strings.Trim(newFanInfo, ",")
	err = UpdateValueByName("user_name", userName, &UserInfo{}, "fans_names", newFanInfo)
	if err != nil {
		return err
	}
	err = UpdateValueByName("user_name", userName, &UserInfo{}, "num_fans", u.NumFans-1)
	if err != nil {
		return err
	}
	return nil
}

// IsFollowedUser 判断该用户是否已关注
func (user *UserInfo) IsFollowedUser(userName string, idolName string) (map[string]interface{}, bool, error) {
	isExist := user.CheckUserExist(userName)
	if !isExist {
		return nil, false, errors.New("该用户不存在")
	}
	isExist = user.CheckUserExist(idolName)
	if !isExist {
		return nil, false, errors.New("关注的用户不存在")
	}
	m, err := user.GetIdolsNames(userName)
	if err != nil {
		return nil, false, err
	}
	_, ok := m[idolName]
	if ok {
		return nil, true, nil
	}
	return m, false, nil
}

// FollowUser 关注或取消关注用户
func (user *UserInfo) FollowUser(request *service.UserRequest2) codeMsg.CodeMessage {
	// choose : 1-关注；0-取消关注
	_, b, err := user.IsFollowedUser(request.UserInfo.UserName, request.UserInfo.IdolsNames)
	if err != nil {
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: err.Error()}
	}
	fmt.Println("b=", b)

	if request.Choose == 1 { // 关注
		if b {
			return codeMsg.CodeMessage{StatusCode: codeMsg.SUCCESS, StatusMessage: "用户已关注"}
		}
		err = user.addIdol(request.UserInfo.UserName, request.UserInfo.IdolsNames) // 当前用户更新关注信息
		if err != nil {
			return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: err.Error()}
		}
		err = user.addFan(request.UserInfo.IdolsNames, request.UserInfo.UserName) // 被关注用户更新粉丝信息
		if err != nil {
			return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: err.Error()}
		}

		return codeMsg.CodeMessage{StatusCode: codeMsg.SUCCESS, StatusMessage: "关注成功"}
	} else { // 取消关注
		if b {
			err = user.deleteIdol(request.UserInfo.UserName, request.UserInfo.IdolsNames) // 当前用户更新关注信息
			if err != nil {
				return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: err.Error()}
			}
			err = user.deleteFan(request.UserInfo.IdolsNames, request.UserInfo.UserName) // 被关注用户更新粉丝信息
			if err != nil {
				return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: err.Error()}
			}
			return codeMsg.CodeMessage{StatusCode: codeMsg.SUCCESS, StatusMessage: "取消关注成功"}
		}
		return codeMsg.CodeMessage{StatusCode: codeMsg.SUCCESS, StatusMessage: "用户未关注"}
	}
}

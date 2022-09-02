package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"wendaxitong/api_gin_gateway/internal/service"
	"wendaxitong/api_gin_gateway/pkg/codeMsg"
	"wendaxitong/api_gin_gateway/pkg/util"
)

type User struct {
	UserId   uint64 `json:"user_id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	var user service.UserRequest
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		c.Abort()
		return
	}
	if user.UserInfo.UserName == "" || user.UserInfo.Password == "" || user.UserInfo.Phone == "" || user.UserInfo.Email == "" {
		c.JSON(http.StatusOK, gin.H{"error": "账号、密码、手机号、邮箱均不能为空"})
		c.Abort()
		return
	}
	request := &user
	response, err := GrpcUerServiceClient.UserRegister(context.Background(), request)
	if err != nil {
		c.JSON(http.StatusOK, "注册失败："+err.Error())
		c.Abort()
		return
	}
	if response.StatusCode != codeMsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code": response.StatusCode,
			"msg":  response.StatusMessage,
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, util.JsonData{
		Code:    response.StatusCode,
		Message: response.StatusMessage,
		Data:    response.UserInfo,
	})
}

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
	var user User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, util.JsonData{
			Code:    codeMsg.Failed,
			Message: err.Error(),
			Data:    "null",
		})
		fmt.Println(err.Error())
		c.Abort()
		return
	}

	if user.UserName == "" || user.Password == "" {
		c.JSON(http.StatusOK, gin.H{"error": "账号、密码均不能为空"})
		c.Abort()
		return
	}

	request := &service.UserRequest{
		UserInfo: &service.UserModel{
			UserName: user.UserName,
			Password: user.Password,
		},
	}
	response, err := GrpcUerServiceClient.UserLogin(context.Background(), request)
	if err != nil {
		fmt.Println("Err:", err.Error())
		c.JSON(http.StatusOK, err.Error())
		c.Abort()
		return
	}
	if response.StatusCode != codeMsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code": response.StatusCode,
			"msg":  response.StatusMessage,
		})
		c.Abort()
		return
	}

	// 生成token
	accessTokenString, _, err := util.GetARToken(user.UserName, user.Password)
	if err != nil {
		c.JSON(http.StatusOK, util.JsonData{
			Code:    codeMsg.Failed,
			Message: err.Error(),
			Data:    "null",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, util.JsonData{
		Code:    response.StatusCode,
		Message: response.StatusMessage,
		Data: gin.H{
			"code":         codeMsg.SUCCESS,
			"access_token": accessTokenString,
			"user_info":    response.UserInfo,
		},
	})
}

// DeleteUser 注销用户信息
func DeleteUser(c *gin.Context) {
	userName := c.MustGet("userName").(string)
	if userName == "" {
		c.JSON(http.StatusOK, "userName为空，参数传递错误")
		c.Abort()
		return
	}

	// 从数据库中物理删除用户信息
	request := &service.UserRequest{
		UserInfo: &service.UserModel{
			UserName: userName,
		},
	}

	response, err := GrpcUerServiceClient.DeleteUser(context.Background(), request)
	if err != nil {
		fmt.Println("GrpcUerServiceClient.DeleteUser Error:", err.Error())
		c.JSON(http.StatusOK, err.Error())
		c.Abort()
		return
	}

	var s = " "
	err1 := util.DeleteARToken(userName) // 用户注销后从redis中删除token信息
	if err1 != nil {
		fmt.Println("DeleteARToken() err: ", err1.Error())
		s = s + err1.Error()
	}

	c.JSON(http.StatusOK, util.JsonData{
		Code:    response.StatusCode,
		Message: response.StatusMessage + s,
		Data:    "null",
	})

}

// UserExit 用户退出登录
func UserExit(c *gin.Context) {
	userName := c.MustGet("userName").(string)
	if userName == "" {
		c.JSON(http.StatusOK, "userName为空，参数传递错误")
		c.Abort()
		return
	}

	err := util.DeleteARToken(userName) // 用户退出登录从redis中删除token信息
	if err != nil {
		fmt.Println("DeleteARToken() err: ", err.Error())
		c.JSON(http.StatusOK, util.JsonData{
			Code:    codeMsg.Failed,
			Message: "退出失败，err:" + err.Error(),
			Data:    "null",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, util.JsonData{
		Code:    codeMsg.SUCCESS,
		Message: "退出登录成功",
		Data:    "null",
	})
}

// ModifyUserInfo 修改个人信息
func ModifyUserInfo(c *gin.Context) {
	userName := c.MustGet("userName").(string)
	if userName == "" {
		c.JSON(http.StatusOK, "userName为空，参数传递错误")
		c.Abort()
		return
	}

	var modifyInfo service.UserRequest
	err := c.ShouldBind(&modifyInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		c.Abort()
		return
	}
	// 根据用户名获取用户id
	request := &service.UserRequest{
		UserInfo: &service.UserModel{
			UserName: userName,
		},
	}
	response, err := GrpcUerServiceClient.GetUserIdByUserName(context.Background(), request)
	if err != nil {
		fmt.Println("GrpcUerServiceClient.GetUserIdByUserName Error:", err.Error())
		c.JSON(http.StatusOK, err.Error())
		c.Abort()
		return
	}

	// 发送修改用户个人信息的请求
	modifyInfo.UserInfo.UserId = response.UserInfo.UserId
	request = &modifyInfo
	response, err = GrpcUerServiceClient.ModifyUserInfo(context.Background(), request)
	if err != nil {
		fmt.Println("GrpcUerServiceClient.ModifyUserInfo Error:", err.Error())
		c.JSON(http.StatusOK, err.Error())
		c.Abort()
		return
	}

	if modifyInfo.UserInfo.UserName != "" || modifyInfo.UserInfo.Password != "" {
		err = util.DeleteARToken(userName) // 用户退出登录从redis中删除token信息
		if err != nil {
			log.Fatalln("token删除异常： ", err.Error())
		}
		c.JSON(http.StatusOK, util.JsonData{
			Code:    response.StatusCode,
			Message: response.StatusMessage + ",请重新登录",
			Data: User{
				UserId:   response.UserInfo.UserId,
				UserName: response.UserInfo.UserName,
				Password: "******",
				Phone:    response.UserInfo.Phone,
				Email:    response.UserInfo.Email,
			},
		})
		return
	}

	c.JSON(http.StatusOK, util.JsonData{
		Code:    response.StatusCode,
		Message: response.StatusMessage,
		Data: User{
			UserId:   response.UserInfo.UserId,
			UserName: response.UserInfo.UserName,
			Password: "******",
			Phone:    response.UserInfo.Phone,
			Email:    response.UserInfo.Email,
		},
	})
}

// GetAccessToken 测试辅助工具，快速获取双token中的accessToken
func GetAccessToken(c *gin.Context) {
	var user User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, util.JsonData{
			Code:    codeMsg.ErrorInvalidParameter,
			Message: "无效参数",
			Data:    "",
		})
		fmt.Println(err.Error())
		c.Abort()
		return
	}

	token, err := util.GetValueByKey(user.UserName + util.AccessTokenKeySuffix)
	if err != nil {
		if err != nil {
			// accessToken过期，刷新accessToken
			var code uint64
			_, code, err = util.RefreshAccessToken(user.UserName)
			if code == codeMsg.Failed {
				c.JSON(http.StatusOK, util.JsonData{
					Code:    codeMsg.Failed,
					Message: "RefreshAccessToken() err:" + err.Error(),
					Data:    "null",
				})
				c.Abort()
				return
			} else if code == codeMsg.ErrorInvalidToken {
				c.JSON(http.StatusOK, util.JsonData{
					Code:    codeMsg.ErrorInvalidToken,
					Message: "token失效，请重新登录",
					Data:    "null",
				})
				c.Abort()
				return
			}
		}
	}
	token, err = util.GetValueByKey(user.UserName + util.AccessTokenKeySuffix)
	if err != nil {
		c.JSON(http.StatusOK, util.JsonData{
			Code:    codeMsg.Failed,
			Message: "重新登录",
			Data:    "",
		})
		fmt.Println(err.Error())
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user_name":    user.UserName,
		"access_token": token,
	})
}

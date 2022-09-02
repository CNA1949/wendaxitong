package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wendaxitong/api_gin_gateway/internal/service"
	"wendaxitong/api_gin_gateway/pkg/codeMsg"
	"wendaxitong/api_gin_gateway/pkg/util"
)

type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	var user service.UserRequest
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	if user.UserName == "" || user.Password == "" {
		c.JSON(http.StatusOK, gin.H{"error": "账号或密码不能为空"})
		return
	}
	request := &service.UserRequest{
		UserName: user.UserName,
		Password: user.Password,
		Phone:    user.Phone,
		Email:    user.Email,
	}
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
		c.JSON(http.StatusOK, gin.H{"code": 2001, "codeMsg": "无效参数"})
		fmt.Println(err.Error())
		return
	}
	request := &service.UserRequest{
		UserName: user.UserName,
		Password: user.Password,
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
	accessTokenString, refreshTokenString, err := util.GetARToken(user.UserName, user.Password)
	if err != nil {
		c.JSON(http.StatusOK, util.JsonData{
			Code:    codeMsg.Failed,
			Message: err.Error(),
			Data:    "null",
		})
	}

	c.JSON(http.StatusOK, util.JsonData{
		Code:    response.StatusCode,
		Message: response.StatusMessage,
		Data: gin.H{
			"token":     accessTokenString + " " + refreshTokenString,
			"user_info": response.UserInfo,
		},
	})
}

// DeleteUser 注销用户信息
func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, util.JsonData{
		Code:    "2002",
		Message: "用户已删除",
		Data:    "null",
	})
}

// UserExit 用户退出登录
func UserExit(c *gin.Context) {
	c.JSON(http.StatusOK, util.JsonData{
		Code:    "2002",
		Message: "用户已退出",
		Data:    "null",
	})
}

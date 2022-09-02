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
		c.Abort()
		return
	}
	if user.UserName == "" || user.Password == "" || user.Phone == "" || user.Email == "" {
		c.JSON(http.StatusOK, gin.H{"error": "账号、密码、手机号、邮箱均不能为空"})
		c.Abort()
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
	token := c.GetHeader("token")
	if token == "" {
		c.JSON(http.StatusOK, util.JsonData{
			Code:    codeMsg.ErrorInvalidParameter,
			Message: "无效参数",
			Data:    "null",
		})
		c.Abort()
		return
	}
	claims, err := util.ParseToken(token, util.AccessSecret)
	if err != nil {
		c.JSON(http.StatusOK, util.JsonData{
			Code:    codeMsg.Failed,
			Message: "Error ParseToken():" + err.Error(),
			Data:    "null",
		})
	}

	// 从数据库中物理删除用户信息
	request := &service.UserRequest{
		UserName: claims.Name,
	}
	response, err := GrpcUerServiceClient.DeleteUser(context.Background(), request)
	if err != nil {
		fmt.Println("GrpcUerServiceClient.DeleteUser Error:", err.Error())
		c.JSON(http.StatusOK, err.Error())
		c.Abort()
		return
	}

	err1 := util.DeleteARToken(claims.Name) // 用户注销后从redis中删除token信息
	if err1 != nil {
		fmt.Println("DeleteARToken() err: ", err1.Error())
	}

	c.JSON(http.StatusOK, util.JsonData{
		Code:    response.StatusCode,
		Message: response.StatusMessage + " " + err1.Error(),
		Data:    "null",
	})

}

// UserExit 用户退出登录
func UserExit(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		c.JSON(http.StatusOK, util.JsonData{
			Code:    codeMsg.ErrorInvalidParameter,
			Message: "无效参数",
			Data:    "null",
		})
		c.Abort()
		return
	}
	claims, err := util.ParseToken(token, util.AccessSecret)
	if err != nil {
		c.JSON(http.StatusOK, util.JsonData{
			Code:    codeMsg.Failed,
			Message: "Error ParseToken():" + err.Error(),
			Data:    "null",
		})
		c.Abort()
		return
	}
	err = util.DeleteARToken(claims.Name) // 用户退出登录从redis中删除token信息
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

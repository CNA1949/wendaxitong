package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wendaxitong/api_gin_gateway/internal/handler"
	"wendaxitong/api_gin_gateway/internal/middleware"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	v1 := ginRouter.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "success")
		})
		user := v1.Group("/user")
		// 用户服务
		user.POST("/login", handler.UserLogin)       // 用户登录
		user.POST("/register", handler.UserRegister) // 用户注册

		// 需要登录保护
		userDo := user.Group("/do")
		userDo.Use(middleware.CheckTokenMiddleware())
		{
			// 用户操作模块
			userDo.POST("/deleteUser", handler.DeleteUser) // 注销账号信息
			userDo.GET("/exit", handler.UserExit)          //退出登录

			follow := userDo.Group("/follow")
			{
				follow.POST("/execute") // 关注与取消关注其他用户
				follow.GET("/list")     // 获取已关注的用户列表
			}

			topic := userDo.Group("/topic")
			{
				topic.GET("/topicList")      // 获取话题列表
				topic.GET("/someoneTopic")   // 获取某一个话题信息
				topic.POST("/commentTopic")  // 评论某个话题
				topic.POST("/deleteTopic")   //删除某个话题
				topic.POST("/deleteComment") //删除某个评论
			}

		}
	}
	return ginRouter
}

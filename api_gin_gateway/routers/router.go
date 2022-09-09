package routers

import (
	"github.com/gin-gonic/gin"
	"wendaxitong/api_gin_gateway/internal/handler"
	"wendaxitong/api_gin_gateway/internal/middleware"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	v1 := ginRouter.Group("/v1")
	{
		user := v1.Group("/user")
		{
			{
				// 用户服务
				user.POST("/login", handler.UserLogin)             // 用户登录
				user.POST("/register", handler.UserRegister)       // 用户注册
				user.POST("/queryUserInfo", handler.QueryUserInfo) // 用户查询某个其他用户的信息
				//测试辅助
				user.POST("/getAccessToken", handler.GetAccessToken)
			}

			userDo := user.Group("/do")
			userDo.Use(middleware.CheckTokenMiddleware())
			{
				// 用户操作模块
				userDo.GET("/deleteUser", handler.DeleteUser)  // 注销账号信息
				userDo.GET("/exit", handler.UserExit)          //退出登录
				userDo.POST("/modify", handler.ModifyUserInfo) // 修改个人信息

				follow := userDo.Group("/follow")
				{
					follow.POST("/following", handler.FollowUser)            // 关注与取消关注其他用户
					follow.GET("/userIdolsList", handler.QueryUserIdolsList) // 查询所有已关注用户
					follow.GET("/userFansList", handler.QueryUserFansList)   // 查询用户拥有的粉丝
				}

				interact := userDo.Group("/interact")
				{
					interact.POST("/createTopic")   // 创建话题
					interact.POST("/deleteTopic")   //删除某个话题
					interact.POST("/someoneTopic")  // 获取话题的具体信息
					interact.POST("/commentTopic")  // 评论某个话题
					interact.POST("/deleteComment") //删除某个评论
					interact.POST("/likesTopic")    //点赞话题
					interact.POST("/likesComment")  // 点赞评论
				}

			}
		}

	}
	return ginRouter
}

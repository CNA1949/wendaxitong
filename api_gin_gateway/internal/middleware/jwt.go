package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wendaxitong/api_gin_gateway/internal/handler"
	"wendaxitong/api_gin_gateway/pkg/codeMsg"
	"wendaxitong/api_gin_gateway/pkg/util"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

////CheckTokenMiddleware 用鉴权到中间件
//func CheckTokenMiddleware() func(c *gin.Context) {
//	return func(c *gin.Context) {
//		token := c.Request.Header.Get("token")
//		tokenStrings := strings.Split(token, " ")
//		if len(tokenStrings) != 2 {
//			c.JSON(http.StatusOK, gin.H{
//				"code":    2003,
//				"codeMsg": "token为空",
//			})
//			c.Abort()
//			return
//		}
//
//		parseToken, isUpd, err := util.ParseToken(tokenStrings[0], tokenStrings[1])
//		if err != nil {
//			c.JSON(http.StatusOK, gin.H{
//				"code":    2005,
//				"codeMsg": "无效的Token",
//			})
//			c.Abort()
//			return
//		}
//		if isUpd {
//			// accessToken已经失效，需要刷新双Token
//			tokenStrings[0], tokenStrings[1] = util.GetToken(parseToken.Name, parseToken.Password)
//		}
//		c.Set("userName", parseToken.Name)
//		c.Next()
//	}
//}

//CheckTokenMiddleware 用鉴权到中间件
func CheckTokenMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		var user handler.User

		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusOK, util.JsonData{
				Code:    codeMsg.ErrorInvalidParameter,
				Message: "无效参数，参数应为UserName或user_name(json)",
				Data:    "null",
			})
			c.Abort()
			return
		}
		//var accessToken string
		_, err = util.GetValueByKey(user.UserName + util.AccessTokenKeySuffix)
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

		c.Set("userName", user.UserName)
		c.Next()
	}
}

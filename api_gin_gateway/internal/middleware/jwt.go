package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wendaxitong/api_gin_gateway/pkg/codeMsg"
	"wendaxitong/api_gin_gateway/pkg/util"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// CheckTokenMiddleware 权限认证
func CheckTokenMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		claims, err := util.ParseToken(token, util.AccessSecret)
		if err != nil {
			c.JSON(http.StatusOK, util.JsonData{
				Code:    codeMsg.Failed,
				Message: "ParseToken() err:" + err.Error(),
				Data:    "null",
			})
			c.Abort()
			return
		}

		// 验证token的有效性
		_, err = util.GetValueByKey(claims.Name + util.AccessTokenKeySuffix)
		if err != nil {
			// accessToken过期，刷新accessToken
			var code uint64
			_, code, err = util.RefreshAccessToken(claims.Name)
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
		token, err = util.GetValueByKey(claims.Name + util.AccessTokenKeySuffix)
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

		c.Set("userName", claims.Name)
		c.Next()
	}
}

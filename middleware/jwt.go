package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zou2699/learnGin2/pkg/app"
	"github.com/zou2699/learnGin2/pkg/e"
	"github.com/zou2699/learnGin2/pkg/utils"
	"net/http"
	"time"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			appG  = app.Gin{C: c}
			code  = e.Success
			token string
		)

		// 从url或者header提取token
		if c.Query("token") != "" {
			token = c.Query("token")
		} else if c.GetHeader("token") != "" {
			token = c.GetHeader("token")
		} else {
			token, _ = c.Cookie("token")
		}

		if token == "" {
			code = e.ErrorAuthCheckTokenFail
			appG.Response(http.StatusUnauthorized, code, nil)

			c.Abort()
			return
		}

		claims, err := utils.ParseToken(token)
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
			appG.Response(http.StatusUnauthorized, code, nil)

			c.Abort()
			return
		}
		if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
			appG.Response(http.StatusUnauthorized, code, nil)

			c.Abort()
			return
		}
		c.Next()
	}

}

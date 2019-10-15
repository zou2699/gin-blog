/*
@Time : 2019/10/15 10:15
@Author : Tux
@File : userinfo
@Description :
*/

package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-blog/pkg/app"
	"gin-blog/pkg/e"
)

func UserInfo(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		code = e.Success
	)

	type UserInfo struct {
		Roles        []string `json:"roles"`
		Introduction string   `json:"introduction"`
		Avatar       string   `json:"avatar"`
		Name         string   `json:"name"`
	}

	var info UserInfo
	info.Roles = []string{"admin"}
	info.Introduction = "I am a super administrator from gin"
	info.Avatar = "https://blog.zouhl.com/favicon-32x32.png"
	info.Name = "Super Admin from gin"

	appG.Response(http.StatusOK, code, info)

}

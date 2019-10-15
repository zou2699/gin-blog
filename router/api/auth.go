package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-blog/model"
	"gin-blog/pkg/app"
	"gin-blog/pkg/e"
	"gin-blog/pkg/utils"
)

type Auth struct {
	Username string `form:"username" json:"username" xml:"username"  binding:"required" valid:"Required;MaxSize(50)"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required" valid:"Required;MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		code = e.Success
		err  error
	)
	var loginForm Auth
	if err := c.ShouldBind(&loginForm); err != nil {
		log.Printf("%+v", loginForm)
		code = e.ErrorAuth
		appG.Response(http.StatusUnauthorized, code, nil)

		return
	}
	// username := c.Request.FormValue("username")
	// password := c.Request.FormValue("password")
	username := loginForm.Username
	password := loginForm.Password

	log.Printf("username: %v,password: %v\n", username, password)
	// log.Println("user:",username,"pass:",password)
	if (username == "") || (password == "") {
		code = e.ErrorAuth
		appG.Response(http.StatusUnauthorized, code, nil)

		return
	}

	// todo validation
	// a := Auth{Username: username, Password: password}

	data := make(map[string]interface{})
	err = model.CheckAuth(username, password)
	if err != nil {
		code = e.ErrorAuth
		appG.Response(http.StatusUnauthorized, code, err.Error())

		return
	}
	token, err := utils.GenerateToken(username, password)
	if err != nil {
		code = e.ErrorAuth
		appG.Response(http.StatusUnauthorized, code, err.Error())

		return
	}
	data["token"] = token
	code = e.AuthSuccess
	// c.SetCookie("token", token, 3600, "/", "", false, true)
	// c.Redirect(http.StatusMovedPermanently, "/api/v1/articles")
	appG.Response(http.StatusOK, code, data)
}

func Logout(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		code = e.Success
	)
	data := make(map[string]interface{})

	data["data"] = "success"
	appG.Response(http.StatusOK, code, data)
}

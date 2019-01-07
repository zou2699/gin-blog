package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zou2699/learnGin2/model"
	"github.com/zou2699/learnGin2/pkg/app"
	"github.com/zou2699/learnGin2/pkg/e"
	"github.com/zou2699/learnGin2/pkg/utils"
	"net/http"
)

type Auth struct {
	Username string `valid:"Required;MaxSize(50)"`
	Password string `valid:"Required;MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		code = e.Success
		err  error
	)

	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	//log.Println("user:",username,"pass:",password)
	if (username == "") || (password == "") {
		code = e.ErrorAuth
		appG.Response(http.StatusUnauthorized, code, nil)

		return
	}

	//todo validation
	//a := Auth{Username: username, Password: password}

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
	c.SetCookie("token", token, 3600, "/", "", false, true)
	c.Redirect(http.StatusMovedPermanently, "/api/v1/articles")
	//appG.Response(http.StatusOK, code, data)
}

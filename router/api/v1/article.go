package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zou2699/learnGin2/model"
	"github.com/zou2699/learnGin2/pkg/app"
	"github.com/zou2699/learnGin2/pkg/e"
	"net/http"
	"strconv"
)

//获取单个文章
func GetArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	code := e.Success

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		code = e.InternalServerError
		appG.Response(http.StatusInternalServerError, code, err.Error())
		return
	}
	article, err := model.GetArticle(id)
	if err != nil {
		code = e.InternalServerError
		appG.Response(http.StatusInternalServerError, code, err.Error())
		return
	}

	appG.Response(http.StatusOK, code, article)
}

//获取多个文章
func GetArticles(c *gin.Context) {
}

//新增文章
func AddArticle(c *gin.Context) {
}

//修改文章
func EditArticle(c *gin.Context) {
}

//删除文章
func DeleteArticle(c *gin.Context) {
}

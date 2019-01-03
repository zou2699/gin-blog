package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zou2699/learnGin2/model"
	"github.com/zou2699/learnGin2/pkg/app"
	"github.com/zou2699/learnGin2/pkg/e"
	"github.com/zou2699/learnGin2/pkg/setting"
	"github.com/zou2699/learnGin2/pkg/utils"
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

	// 判断是否存在
	exist, err := model.ExistArticleByID(id)
	if err != nil {
		code = e.InternalServerError
		appG.Response(http.StatusInternalServerError, code, err.Error())
		return
	}
	if !exist {
		code = e.ErrorNotExistArticle
		appG.Response(http.StatusNotFound, code, nil)
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
	appG := app.Gin{C: c}
	code := e.Success
	var err error

	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	data["list"], err = model.GetArticles(utils.Getpage(c), setting.Server.PageSize, maps)
	if err != nil {
		code = e.InternalServerError
		appG.Response(http.StatusInternalServerError, code, err.Error())
		return
	}

	data["total"], err = model.GetArticleTotal(maps)
	if err != nil {
		code = e.InternalServerError
		appG.Response(http.StatusInternalServerError, code, err.Error())
		return
	}

	appG.Response(http.StatusOK, code, data)
}

//新增文章
func AddArticle(c *gin.Context) {
	var (
		appG    = app.Gin{C: c}
		code    = e.Success
		err     error
		article model.Article
		//data map[string]interface{}
	)

	err = c.ShouldBind(&article)
	//log.Printf("%+v\n", article)
	if err != nil {
		code = e.InternalServerError
		appG.Response(http.StatusInternalServerError, code, err.Error())
		return
	}

	exist, err := model.ExistTagByID(article.TagID)
	if err != nil {
		code = e.InternalServerError
		appG.Response(http.StatusInternalServerError, code, err.Error())
		return
	}
	if !exist {
		code = e.ErrorNotExistTag
		appG.Response(http.StatusNotFound, code, nil)
		return
	}

	//data["tag_id"] = article.TagID
	//	data["title"] = article.Title
	//	data["desc"]= article.Desc
	//	data["content"] = article.Content
	//	data["create_by"] = article.CreateBy
	//	data["state"]=article.State

	err = model.AddArticle(article)
	if err != nil {
		code = e.InternalServerError
		appG.Response(http.StatusInternalServerError, code, err.Error())
		return
	}

	appG.Response(http.StatusOK, code, nil)
}

//修改文章
func EditArticle(c *gin.Context) {
	var (
		appG    = app.Gin{C: c}
		code    = e.Success
		err     error
		article model.Article
		//data map[string]interface{}
	)

	err = c.ShouldBind(&article)
	if err != nil {
		code = e.InternalServerError
		appG.Response(http.StatusInternalServerError, code, err.Error())
		return
	}

	//获取文章id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		code = e.InternalServerError
		appG.Response(http.StatusInternalServerError, code, err.Error())

		return
	}
	article.ID = id
	//log.Printf("%+v\n", article)

	// 文章是否存在
	exist, err := model.ExistArticleByID(article.ID)
	if err != nil {
		code = e.InternalServerError
		appG.Response(http.StatusInternalServerError, code, err.Error())
		return
	}
	if !exist {
		code = e.ErrorNotExistArticle
		appG.Response(http.StatusNotFound, code, nil)
		return
	}
	// 新标签id是否存在
	existTag, err := model.ExistTagByID(article.TagID)
	if err != nil {
		code = e.InternalServerError
		appG.Response(http.StatusInternalServerError, code, err.Error())
		return
	}
	if !existTag {
		code = e.ErrorNotExistTag
		appG.Response(http.StatusNotFound, code, nil)
		return
	}

	err = model.EditArticle(article)
	if err != nil {
		code = e.InternalServerError
		appG.Response(http.StatusInternalServerError, code, err.Error())
		return
	}

	appG.Response(http.StatusOK, code, nil)
}

//删除文章
func DeleteArticle(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		code = e.Success
		err  error
	)

	//获取文章id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		code = e.InternalServerError
		appG.Response(http.StatusInternalServerError, code, err.Error())

		return
	}

	exist, err := model.ExistArticleByID(id)
	if err != nil {
		code = e.InternalServerError
		appG.Response(http.StatusInternalServerError, code, err.Error())
		return
	}
	if !exist {
		code = e.ErrorNotExistArticle
		appG.Response(http.StatusNotFound, code, nil)
		return
	}

	err = model.DeleteArticle(id)
	if err != nil {
		code = e.InternalServerError
		appG.Response(http.StatusInternalServerError, code, err.Error())
		return
	}
	appG.Response(http.StatusOK, code, nil)

}

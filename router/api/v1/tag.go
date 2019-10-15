package v1

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-blog/model"
	"gin-blog/pkg/app"
	"gin-blog/pkg/e"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/utils"
)

// 获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state = -1
	if arg := c.Query("state"); arg != "" {
		state, _ := strconv.Atoi(arg)
		maps["state"] = state
	} else {
		maps["state"] = state
	}
	code := e.Success

	data["lists"] = model.GetTags(utils.Getpage(c), setting.Server.PageSize, maps)
	data["total"] = model.GetTagTotal(maps)

	code = 20001
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 新增文章标签
func AddTag(c *gin.Context) {
	appG := app.Gin{C: c}
	code := e.Success
	var tag model.Tag

	err := c.ShouldBind(&tag)
	if err != nil {
		code = e.InternalServerError
		log.Println("bind tag Error of Addtag ", err.Error())
		appG.Response(http.StatusInternalServerError, code, err.Error())
		return
	}

	// todo validation

	model.AddTag(tag.Name, tag.State, tag.CreatedBy)
	appG.Response(http.StatusOK, code, nil)
}

// 编辑文章标签
func EditTag(c *gin.Context) {
	appG := app.Gin{C: c}
	code := e.Success
	var tag model.Tag

	err := c.ShouldBind(&tag)
	// set delete to nil
	// tag.DeletedAt = nil

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		code = e.InternalServerError
		appG.Response(http.StatusInternalServerError, code, err.Error())

		return
	}
	tag.ID = id

	if err != nil {
		code = e.InternalServerError
		appG.Response(http.StatusInternalServerError, code, err.Error())

		log.Println("bind tag Error of EditTag", err.Error())

		return
	}

	// todo validation

	exists, err := model.ExistTagByID(tag.ID)

	if err != nil {
		code = e.InternalServerError
		appG.Response(http.StatusInternalServerError, code, err.Error())

		log.Println("EditTag Error", err.Error())

		return
	}

	if !exists {
		code = e.ErrorNotExistTag
		appG.Response(http.StatusNotFound, code, nil)

		return
	}

	err = model.EditTag(tag.ID, tag)
	if err != nil {
		code = e.InternalServerError
		appG.Response(http.StatusInternalServerError, code, err.Error())

		return
	}

	appG.Response(http.StatusOK, code, nil)
}

// 删除文章标签
func DeleteTag(c *gin.Context) {
	appG := app.Gin{C: c}
	code := e.Success

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		code = e.InternalServerError
		appG.Response(http.StatusInternalServerError, code, err.Error())

		return
	}

	exist, err := model.ExistTagByID(id)
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

	err = model.DeleteTag(id)
	if err != nil {
		code = e.InternalServerError
		appG.Response(http.StatusInternalServerError, code, err.Error())

		return
	}

	appG.Response(http.StatusOK, code, nil)
}

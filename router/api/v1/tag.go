package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zou2699/learnGin2/model"
	"github.com/zou2699/learnGin2/pkg/e"
	"github.com/zou2699/learnGin2/pkg/setting"
	"github.com/zou2699/learnGin2/pkg/utils"
	"log"
	"net/http"
	"strconv"
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

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 新增文章标签
func AddTag(c *gin.Context) {
	var tag model.Tag
	err := c.ShouldBind(&tag)
	if err != nil {
		log.Println("bind tag Error of Addtag ", err.Error())
	}

	// todo validation
	code := e.Success
	model.AddTag(tag.Name, tag.State, tag.CreatedBy)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

// 编辑文章标签
func EditTag(c *gin.Context) {
	var tag model.Tag
	err := c.ShouldBind(&tag)
	// set delete to nil
	tag.DeletedAt = nil
	log.Println(tag)
	if err != nil {
		log.Println("bind tag Error of EditTag", err.Error())
	}

	// todo validation

	code := e.Success
	exists, err := model.ExistTagByID(tag.ID)
	log.Println(tag.ID, exists, err)
	if err != nil {
		log.Println("EditTag Error", err.Error())
		code = e.InternalServerError
	}
	if !exists {
		code = e.ErrorNotExistTag
	}

	err = model.EditTag(tag.ID, tag)
	if err != nil {
		code = e.InternalServerError
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

// 删除文章标签
func DeleteTag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	code := e.Success
	if err != nil {
		code = e.InternalServerError
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": err.Error(),
		})
		return
	}

	exist, err := model.ExistTagByID(id)
	if err != nil {
		code = e.InternalServerError

		return
	}
	if !exist {
		code = e.ErrorNotExistTag
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": err.Error(),
		})
		return
	}
	err = model.DeleteTag(id)
	if err != nil {
		code = e.InternalServerError
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}

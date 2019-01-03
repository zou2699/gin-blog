package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zou2699/learnGin2/middleware"
	"github.com/zou2699/learnGin2/router/api"
	"github.com/zou2699/learnGin2/router/api/v1"
)

func InitRouter() *gin.Engine {
	/*
		    r.Use(gin.Logger())

		    r.Use(gin.Recovery())
			// env
		    gin.SetMode(setting.RunMode)
	*/
	r := gin.Default()
	r.GET("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.Jwt())
	{
		// tag
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		// article
		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}

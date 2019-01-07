package static

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Loginhtml(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{"title": "Login Index"})
}

func Indexhtml(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "index"})
}

func Bloghtml(c *gin.Context) {

	c.HTML(http.StatusOK, "blog.tmpl", gin.H{
		"title": "blog",
	})
}

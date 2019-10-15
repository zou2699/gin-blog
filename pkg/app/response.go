package app

import (
	"github.com/gin-gonic/gin"

	"gin-blog/pkg/e"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, code int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

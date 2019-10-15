package utils

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-blog/pkg/setting"
)

func Getpage(c *gin.Context) int {
	result := 0

	if c.Query("page") == "" {
		return result
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		log.Println("Query page ERROR", err.Error())
		return result
	}
	if page > 0 {
		result = (page - 1) * setting.Server.PageSize
	}
	return result

}

package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/zou2699/learnGin2/pkg/setting"
	"log"
	"strconv"
)

func Getpage(c *gin.Context) int {
	result := 0

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

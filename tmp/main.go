/*
@Time : 2019/10/14 14:56
@Author : Tux
@File : main
@Description :
*/

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-blog/middleware"
)

type Login struct {
	UserName string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	router.Use(middleware.Cors())
	// Example for binding JSON ({"user": "manu", "password": "123"})
	router.POST("/auth", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.UserName != "admin" || json.Password != "123321" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Example for binding XML (
	//	<?xml version="1.0" encoding="UTF-8"?>
	//	<root>
	//		<user>user</user>
	//		<password>123</password>
	//	</root>)

	// Example for binding a HTML form (user=manu&password=123)
	router.POST("/auth1", func(c *gin.Context) {
		var form Login
		// This will infer what binder to use depending on the content-type header.
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Printf("username %v, password %v\n", form.UserName, form.Password)
		if form.UserName != "admin" || form.Password != "123123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8000")
}

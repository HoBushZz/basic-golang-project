package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	//绑定JSON
	r.POST("/loginJSON", func(ctx *gin.Context) {
		var login Login
		if err := ctx.ShouldBind(&login); err == nil {
			fmt.Printf("login info: %#v\n", login)
			ctx.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}
	})
	//绑定QueryString,虽然查询字符串的数据并不是以表单的形式传递，但在Gin框架中，查询字符串的数据也可以通过form标签进行绑定。
	r.GET("/loginForm", func(ctx *gin.Context) {
		var login Login

		if err := ctx.ShouldBind(&login); err == nil {
			fmt.Printf("login info: %#v\n", login)
			ctx.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}
	})
	//绑定FORM表单
	r.POST("/loginForm", func(ctx *gin.Context) {
		var login Login

		if err := ctx.ShouldBind(&login); err == nil {
			fmt.Printf("login info: %#v\n", login)
			ctx.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}
	})
	r.Run(":9090")
}

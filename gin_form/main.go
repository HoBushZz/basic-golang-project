package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(ctx *gin.Context) {
		/* 获取form表单数据的第一种方法 */
		/* username := ctx.PostForm("username")
		password := ctx.PostForm("password") */
		/* 第二种方法 */
		/* username := ctx.DefaultPostForm("username", "somebody")
		password := ctx.DefaultPostForm("xxx", "***")//有password会当成一个空字符串 */
		/* 第三种方法 */
		username, ok := ctx.GetPostForm("username")
		if !ok {
			username = "somebody"
		}
		password, ok := ctx.GetPostForm("password")
		if !ok {
			password = "***"
		}
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})

	r.Run(":9090")
}

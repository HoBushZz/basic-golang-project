package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	/* HTTP重定向 */
	r.GET("/index", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	/* 路由重定向 */
	r.GET("/a", func(ctx *gin.Context) {
		//跳转到/b对应的路由处理函数
		ctx.Request.URL.Path = "/b" //把请求的URI修改
		r.HandleContext(ctx)
	})
	r.GET("/b", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
	r.Run(":9090")
}

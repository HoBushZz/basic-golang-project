package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* 获取请求的URI参数 */
func main() {
	r := gin.Default()

	r.GET("/:name/:age", func(ctx *gin.Context) {
		name := ctx.Param("name")
		age := ctx.Param("age")
		ctx.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.Run(":9090")
}

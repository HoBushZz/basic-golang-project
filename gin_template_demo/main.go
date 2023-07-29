package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/statics", "./statics") /* 给./statics文件夹的标注为静态文件并在路由前缀加上/statics */
	r.LoadHTMLFiles("./templates/index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil) /* 未define的情况模板名等于文件名，不会带有路径 */
	})
	r.Run(":8090")
}

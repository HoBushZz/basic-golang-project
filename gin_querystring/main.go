package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {
		//	获取浏览器发送请求时携带的queryString数据
		name := c.Query("query") //通过c.Query获得参数
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})
	r.GET("/web2", func(c *gin.Context) {
		name := c.DefaultQuery("query", "somebody")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})
	r.GET("/web3", func(c *gin.Context) {
		name, ok := c.GetQuery("query")
		if !ok {
			name = "somebody"
		}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})
	r.Run(":9090")
}

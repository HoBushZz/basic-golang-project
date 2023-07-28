package main

import (
	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello golang!",
	})
}
func main() {
	r := gin.Default() //返回默认的路由引擎

	r.GET("/hello", sayHello) /*指定用户使用Get请求时，执行sayHello这个函数*/

	r.Run(":9090") /*启动服务*/
}

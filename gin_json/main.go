package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		/*方法一：使用map渲染json*/
		/*data := map[string]interface{}{
			"name":    "Hobush",
			"message": "hello world",
			"age":     22,
		}*/
		/*方法二：gin.H是gin框架提供的map[string]interface{}的一个快捷方式*/
		data2 := gin.H{
			"name":    "Hobush",
			"message": "hello world",
			"age":     22,
		}
		c.JSON(http.StatusOK, data2)
	})
	/*可以使用结构体json-tag规定json中Key值*/
	type msg struct {
		Name    string `json:"name"`
		Message string `json:"msg"`
		Age     int    `json:"age"`
	}
	r.GET("/json_struct", func(c *gin.Context) {
		data := msg{
			"Hobush_struct", "hello golang", 22,
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run(":9090")
}

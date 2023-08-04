package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/* Handlerfunc */
func indexHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("index")
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "index",
		})
	}
}

/* 定义一个中间件m1：统计请求处理函数的耗时 */
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	start := time.Now()
	c.Set("name", "Hobush") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值

	c.Next()
	// c.Abort() //Abort常用于测试连接失败数据库、登录失败后的处理
	// 如果m1后续都不想执行，return
	cost := time.Since(start)
	fmt.Printf("cost:%#v\n", cost)
	fmt.Println("m1 out ...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in ...")
	// go funcxx(c.copy())开启新的协程只能使用副本上下文
	name, ok := c.Get("name") /* 接收上一个组件传来的信息 */
	if !ok {
		name = "匿名用户"
	}
	fmt.Printf("name:#%v\n", name)
	c.Next()
	fmt.Println("m2 out ...")
}
func authMiddleware(doCheck bool) gin.HandlerFunc {
	//连接数据库或者其他一些准备工作
	return func(ctx *gin.Context) {
		if doCheck {
			//存放具体逻辑
			// 是否登录的判断
			// 是否为登录用户
			// c.Next()
			// else c.Abort()
		} else {
			ctx.Next()
		}
	}
}

func main() {
	r := gin.Default() /* 默认使用了logger和Recovery中间件 */
	// r := gin.New() 不使用任何中间件
	r.Use(m1, m2) /* 全局注册中间件函数m1,m2 */
	r.GET("/index", indexHandler())
	r.GET("/shop", func(ctx *gin.Context) {
		fmt.Println("shop")
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})

	/* 给路由组注册中间组件 */
	xxGroup := r.Group("/xx", authMiddleware(true))
	{
		xxGroup.GET("/index", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "index",
			})
		})
	}
	r.Run(":9091")
}

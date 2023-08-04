package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/upload", func(ctx *gin.Context) {
		//从请求中读取读取文件
		fileObj, err := ctx.FormFile("f1")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		} else {
			dst := fmt.Sprintf("./%s", fileObj.Filename)
			//dst := path.Join("./", fileObj.Filename)
			ctx.SaveUploadedFile(fileObj, dst)
			ctx.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		}
		//将读取后的文件保存在服务器本地
	})
	r.Run(":9091")
}

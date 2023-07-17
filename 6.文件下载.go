package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()

	route.GET("/download", func(context *gin.Context) {
		context.Header("Content-type", "application/octet-stream")
		// 指定下载的文件的名字
		context.Header("Content-Disposition", "attachment; filename="+"小胖.png")
		// 前后端分离 则不需要以上两行代码
		// 直接响应资源即可
		context.Header("fileName", "xxx.png")
		context.File("upload/Snipaste_2023-07-14_15-49-54.png")
	})

	route.Run(":9797")
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.POST("/upload", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		fmt.Println(file.Size)
		fmt.Println(file.Filename)

		// 保存到指定目录
		context.SaveUploadedFile(file, "./upload/1.png")
		context.JSON(200, gin.H{"msg": "上传成功"})
	})

	route.POST("/uploads", func(context *gin.Context) {
		form, _ := context.MultipartForm()
		files := form.File["uploads"]
		for _, file := range files {
			context.SaveUploadedFile(file, "./upload/"+file.Filename)
		}
		context.JSON(200, gin.H{"msg": fmt.Sprintf("上传成功%d个文件", len(files))})
	})

	route.Run(":8181")
}

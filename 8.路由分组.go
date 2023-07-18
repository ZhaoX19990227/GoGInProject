package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()
	//user := route.Group("/api")
	//
	//user.GET("/user", func(context *gin.Context) {
	//	context.JSON(200, gin.H{"msg": "请求成功"})
	//})

	// 或者写成这样
	user := route.Group("/api")
	{
		user.GET("/user", func(context *gin.Context) {
			context.JSON(200, gin.H{"msg": "请求成功"})
		})
	}

	route.Run(":5959")
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func m1(context *gin.Context) {
	fmt.Println("m1...")
	context.Abort() // 阻止后续执行

	context.Next() // Next之前的是中间件请求，之后的是中间件响应。遇到直接执行下一个中间件的请求 类似递归进入和返回
}

func main() {
	route := gin.Default()

	route.GET("/", m1, func(context *gin.Context) {
		fmt.Println("...")
	})

	route.Run(":6969")
}

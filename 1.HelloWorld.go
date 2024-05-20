package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func _json(c *gin.Context) {

	type UserInfo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		// 不参与序列化 整个password都不会显示
		Password string `json:"-"`
	}

	user := UserInfo{
		"赵翔",
		24,
		"123456",
	}

	c.JSON(200, user)
}

func main() {
	route := gin.Default()
	route.GET("/index", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World!")
	})

	route.GET("/json", _json)

	err := route.Run(":7777")
	if err != nil {
		return
	}
	// 本质就是对 http.ListenAndServe(":7777",route) 的进一步封装
}

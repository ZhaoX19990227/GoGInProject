package main

import (
	"github.com/gin-gonic/gin"
)

func _redirect(c *gin.Context) {
	// 301 永久 302 临时
	c.Redirect(301, "http://www.baidu.com")
}

func main() {
	route := gin.Default()
	route.GET("/baidu", _redirect)

	route.Run(":9292")
}

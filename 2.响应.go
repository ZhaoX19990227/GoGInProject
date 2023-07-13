package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func _xml(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{"user": "小胖", "age": 12, "status": "ok"})
}

func _ymal(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{"user": "小胖", "age": 12, "status": "ok"})
}

func _html(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"user": "小胖", "age": 12, "status": "ok"})
}

func main() {
	route := gin.Default()
	// go中没有相对当前文件的路径，只有相对项目的路径
	// 第一个参数：访问路径，第二个参数文件路径
	// 静态资源
	route.StaticFile("image", "./static/1.jpeg")
	// 静态目录
	route.StaticFS("static", http.Dir("static/static/1.txt"))
	route.LoadHTMLGlob("templates/*")
	route.GET("/xml", _xml)
	route.GET("/yaml", _ymal)
	route.GET("/html", _html)
	route.Run(":9191")
}

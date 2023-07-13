package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func _query(c *gin.Context) {
	// 获取参数为user的值
	// http://localhost:9393/query?user=小胖&user=小猪
	user := c.Query("user")
	fmt.Println(user)                 //小胖
	fmt.Println(c.GetQuery("user"))   //小胖 true
	fmt.Println(c.QueryArray("user")) //[小胖 小猪]

}

/*
*
表单请求
*/
func _form(c *gin.Context) {
	// 获取表单中的name
	fmt.Println(c.PostForm("name"))
}

func main() {
	route := gin.Default()
	route.GET("/query", _query)
	route.POST("/form", _form)
	route.Run(":9393")
}

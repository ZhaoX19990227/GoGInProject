package main

import (
	"GinProject/geeCache"
	"net/http"
)

func main() {
	r := geeCache.New()
	r.GET("/", func(c *geeCache.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *geeCache.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *geeCache.Context) {
		c.JSON(http.StatusOK, geeCache.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}

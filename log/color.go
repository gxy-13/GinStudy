package main

import "github.com/gin-gonic/gin"

func main() {
	// 强制日志颜色化
	gin.ForceConsoleColor()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}

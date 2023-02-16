package main

// 在中间件或handler中启动新的goroutine时，不能使用原始的上下文，必须使用只读副本

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		// 创建在goroutine中使用的副本
		ccp := c.Copy()
		go func() {
			// 用time.Sleep() 模拟一个长任务
			time.Sleep(5 * time.Second)

			// 请注意您使用的是复制的上下文 ccp
			log.Println("Done! in path" + ccp.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {
		// 用time.Sleep() 模拟一个长任务
		time.Sleep(5 * time.Second)
		// 因为没有使用goroutine， 不需要拷贝上下文
		log.Println("Done in path" + c.Request.URL.Path)
	})
	r.Run(":8080")
}

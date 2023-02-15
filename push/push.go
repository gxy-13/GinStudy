package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

var html = template.Must(template.New("https").Parse(
	`
	<html>
	<head>
	  <title>Https Test</title>
	  <script src="/assets/app.js"></script>
	</head>
	<body>
	  <h1 style="color:red;">Welcome, Ginner!</h1>
	</body>
	</html>
`))

// push 还没有收到浏览器请求，服务器就把各种资源推送给浏览器，比如浏览器只请求了index.html，但是把css都发送给浏览器，提高性能

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html)

	r.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// 使用pusher.Push() 做服务器推送
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to Push:%v", err)
			}
		}
		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})
	// 监听并启动
	r.Run()
}

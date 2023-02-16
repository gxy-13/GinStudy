package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"'`
}

func SomeHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}

	// c.ShouldBind 使用了 c.Request.Body 不能重用
	// 想要多次绑定，可以使用c.ShouldBindBodyWith
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// 复用存储在上下文中的body
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		c.String(http.StatusOK, `the body should be formB JSON`)
	}
	// c.ShouldBindBodyWith 会在绑定之前将body存储到上下文，这会对性能造成轻微影响
	// 只有某些格式需要此功能，json xml msgpack，protobuf。其他格式，如Query，Form 可以多次调用
	// c.ShouldBind
}

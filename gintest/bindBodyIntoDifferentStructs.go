/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/3/6
 * Time 2:47 下午
 */

package main

import (
	"net/http"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

type formA struct {
	Foo string `form:"foo" json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `form:"bar" json:"bar" xml:"bar" binding:"required"`
}

func SomeHandler(c *gin.Context) {
	var objA formA
	var objB formB
	// This c.ShouldBind consumes c.Request.Body and it cannot be reused.
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// Always an error is occurred by this because c.Request.Body is EOF now.
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	} else {
		c.String(http.StatusOK, `the body should be nothing`)
	}
}

func main() {
	r := gin.Default()
	r.POST("/body", SomeHandler)
	r.Run()
}

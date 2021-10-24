/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/4/26
 * Time 下午1:49
 */

package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"

	"github.com/gin-gonic/gin"
)

func main() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.New()
	r.Use(BreakerWrapper)
	r.Use(gin.Recovery())
	v1 := r.Group("/v1")
	{
		v1.GET("/panic", Panic)
	}
	return r
}

func Panic(c *gin.Context) {
	c.JSON(200, gin.H{"message": "this is panic!"})
}

func BreakerWrapper(c *gin.Context) {
	name := c.Request.Method + "-" + c.Request.RequestURI
	hystrix.ConfigureCommand(name, hystrix.CommandConfig{
		Timeout:               1000,
		MaxConcurrentRequests: 100,
		ErrorPercentThreshold: 1,
	})
	hystrix.Do(name, func() error {
		c.Next()
		statusCode := c.Writer.Status()
		if statusCode >= http.StatusInternalServerError {
			str := fmt.Sprintf("status code %d", statusCode)
			return errors.New(str)
		}
		return nil
	}, func(err error) error {
		if err == hystrix.ErrCircuitOpen {
			c.String(http.StatusAccepted, "请稍后重试")
		}
		return err
	})
}

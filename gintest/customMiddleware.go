/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/3/5
 * Time 6:07 下午
 */

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin.New()
	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// it would print: "12345"
		log.Println(example)
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
		})
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

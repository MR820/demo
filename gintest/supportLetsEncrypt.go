/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/3/6
 * Time 11:24 上午
 */

package main

import (
	"log"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	log.Fatal(autotls.Run(r, "localhost", "127.0.0.1"))
}

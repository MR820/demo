/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/10/13
 * Time 3:52 下午
 */

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})
	r.Run()
}

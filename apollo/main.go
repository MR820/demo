/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/4/26
 * Time 下午2:51
 */

package main

import (
	"github.com/gin-gonic/gin"
	"imooc.com/ccmouse/learngo/apollo/apollo"
)

func main() {

	r := gin.Default()
	client := apollo.GetClient()


	r.GET("/", func(c *gin.Context) {
		cache := client.GetConfigCache("application")
		val,_ := cache.Get("host")
		c.JSON(200, gin.H{"message": val})
	})
	r.Run()
}

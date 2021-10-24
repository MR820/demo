/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/1
 * Time 下午4:38
 */

package main

import (
	"github.com/gin-gonic/gin"
	"imooc.com/ccmouse/learngo/grpc/tutorial"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/protobuf", func(c *gin.Context) {
		p :=  &tutorial.Person{
			Name:        "xzw",
			Id:          1,
			Email:       "jsjxzw@163.com",
			Phones: []*tutorial.Person_PhoneNumber{
				{
					Number:"15012345678",
					Type:tutorial.Person_MOBILE,
				},
			},
		}
		c.ProtoBuf(http.StatusOK, p)
	})
	r.Run(":8080")
}



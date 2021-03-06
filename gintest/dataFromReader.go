/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/3/5
 * Time 2:26 下午
 */

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("http://mioa.yundasys.com:17140/ioa/public/upload/5dd4e247efa11.jpg")
		if err != nil || response.StatusCode != http.StatusOK {
			panic(err)
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
	router.Run(":8080")
}

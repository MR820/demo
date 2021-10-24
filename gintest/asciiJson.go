/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/3/3
 * Time 5:00 下午
 */

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		data := gin.H{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		c.JSON(http.StatusOK, data)
	})

	r.GET("/securejson", func(c *gin.Context) {
		data := gin.H{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		c.SecureJSON(http.StatusOK, data)
	})

	r.GET("/jsonp", func(c *gin.Context) {
		data := gin.H{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		c.JSONP(http.StatusOK, data)
	})

	r.GET("/asciijson", func(c *gin.Context) {
		data := gin.H{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		c.AsciiJSON(http.StatusOK, data)
	})

	r.GET("/purejson", func(c *gin.Context) {
		data := gin.H{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		c.PureJSON(http.StatusOK, data)
	})

	r.Run()
}

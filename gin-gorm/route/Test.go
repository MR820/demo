/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/12/22
 * Time 4:18 下午
 */

package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(this *gin.RouterGroup) {
	this.GET("/living", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "服务正常"})
	})
}

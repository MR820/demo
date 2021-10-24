/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/5
 * Time 下午1:56
 */

package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetList(c *gin.Context) {
	fmt.Println(c.Get("serviceIp"))
	fmt.Println(c.Get("servicePort"))
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
}

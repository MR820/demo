/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/12/22
 * Time 10:19 上午
 */

package route

import (
	"github.com/gin-gonic/gin"
	"imooc.com/ccmouse/learngo/gin-gorm/user"
)

func User(this *gin.RouterGroup) {
	this.POST("/users", user.CreateUser)
	this.GET("/users", user.FetchAllUser)
	this.GET("/users/:id", user.FetchSingleUser)
	this.PUT("/users/:id", user.UpdateUser)
	this.DELETE("/users/:id", user.DeleteUser)
}

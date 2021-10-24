/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2021/8/4
 * Time 下午5:43
 */

package main

import (
	"net/http"

	"imooc.com/ccmouse/learngo/nacos/nacos"

	"imooc.com/ccmouse/learngo/nacos/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{"data": c.Value})
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "ok"})
	})

	u := r.Group("/user", nacos.SelectOneHealthyInstance)
	{
		u.GET("", user.GetList)
	}

	r.GET("/registerInstance", nacos.RegisterInstance)
	r.GET("deregisterInstance", nacos.DeregisterInstance)
	r.GET("/getService", nacos.GetService)
	r.GET("/selectAllInstances", nacos.SelectAllInstances)
	r.GET("/selectInstances", nacos.SelectInstances)
	//r.GET("/selectOneHealthyInstance", nacos.SelectOneHealthyInstance)
	r.GET("/subscribe", nacos.Subscribe)
	r.GET("/unsubscribe", nacos.Unsubscribe)
	r.GET("/getAllServicesInfo", nacos.GetAllServicesInfo)

	r.Run(":8080")
}

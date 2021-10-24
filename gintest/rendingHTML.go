/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/3/5
 * Time 2:38 下午
 */

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//func main() {
//	r := gin.Default()
//	r.LoadHTMLGlob("gintest/templates/*")
//	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
//	r.GET("index", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "index.tmpl", gin.H{
//			"title":   "Main website",
//			"content": "abcd",
//		})
//	})
//	r.Run()
//}

func main() {
	r := gin.Default()
	//定界符
	r.Delims("{[{", "}]}")

	r.LoadHTMLGlob("gintest/templates/**/*")

	//html := template.Must(template.ParseFiles("gintest/templates/posts/index.tmpl", "gintest/templates/users/index.tmpl"))
	//r.SetHTMLTemplate(html)
	//r.LoadHTMLGlob("gintest/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})

	r.Run()
}

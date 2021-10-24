package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.JSON(http.StatusOK, gin.H{
			"name":   name,
			"action": action,
		})
	})
	r.GET("/send", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"send": "send",
		})
	})
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})
	r.POST("/user/:name/*action", func(c *gin.Context) {
		//c.FullPath() == "/user/:name/*action"
		c.JSON(http.StatusOK, gin.H{
			"path": c.FullPath(),
		})
	})
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")
		c.JSON(http.StatusOK, gin.H{
			"firstname": firstname,
			"lastname":  lastname,
		})
	})
	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")
		c.JSON(http.StatusOK, gin.H{
			"message": message,
			"nick":    nick,
		})
	})
	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.DefaultPostForm("message", "memo")
		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"page":    page,
			"name":    name,
			"message": message,
		})
	})
	r.POST("/postss", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")
		fmt.Printf("ids: %v;names:%v", ids, names)
		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		c.SaveUploadedFile(file, "test/ymzl.txt")
		c.JSON(http.StatusOK, gin.H{
			"status": 2000,
		})
	})
	r.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for i, file := range files {
			fmt.Println(i)
			log.Println(file.Filename)
		}

		c.JSON(http.StatusOK, gin.H{
			"ststus": 200,
		})
	})
	v1 := r.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  200,
				"message": "login",
			})
		})
		v1.POST("/submit", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "submit",
			})
		})
	}
	r.Run()
}

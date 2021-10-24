package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	router.GET("/static", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"static":   "./assets",
			"staticFS": http.Dir("my_file_system"),
		})
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}

/**
 * Createby GoLand
 * User xzw jsjxzw@163.com
 * Date 2020/3/5
 * Time 3:17 下午
 */

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d/%02d", year, month, day)
}

func main() {
	router := gin.Default()
	router.Delims("{[{", "}]}")
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLFiles("gintest/templates/raw.tmpl")

	router.GET("/raw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "raw.tmpl", gin.H{
			"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
		})
	})

	router.Run(":8080")
}

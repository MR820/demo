package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Person2 struct {
	Name       string    `form:"name"`
	Address    string    `form:"address" binding:"required"`
	Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	CreateTime time.Time `form:"createTime" time_format:"unixNano"`
	UnixTime   time.Time `form:"unixTime" time_format:"unix"`
}

func main() {
	r := gin.Default()
	r.GET("/testing", startPage2)
	r.Run()
}

func startPage2(c *gin.Context) {
	var person Person2
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if err := c.ShouldBind(&person); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	log.Println(person.Name)
	log.Println(person.Address)
	log.Println(person.Birthday)
	log.Println(person.CreateTime)
	log.Println(person.UnixTime)

	c.String(http.StatusOK, "Success")
}

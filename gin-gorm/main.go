package main

import (
	"github.com/gin-gonic/gin"
	"imooc.com/ccmouse/learngo/gin-gorm/route"
)

func main() {
	router()
}

func router() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	route.User(v1)
	route.Test(v1)
	router.Run()
}

package main

//File  : main.go
//Author: Simon
//Describe: use the goeureka
//Date  : 2020/12/3

import (
	"github.com/SimonWang00/goeureka"
	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()
	r.GET("hello", func(c *gin.Context) {
		c.String(200, "hello goeureka")
	})
	goeureka.RegisterClient("http://127.0.0.1:8761","myapp", "8080", "43")
	r.Run("127.0.0.1:8000")
}
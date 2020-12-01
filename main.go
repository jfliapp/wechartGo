package main

import (
	"wechartTest/controller"
	"wechartTest/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/login", controller.Login)
	r.Use(utils.MiddleTokenParse)
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "ha  ha ~~~")
	})
	r.Run()
}

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"params":  c.Query("1"),
		})
	})
	r.GET("/qE", qE)
	r.GET("/insert", i)
	r.GET("/select", s)
	r.Run(":9000")

}

package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting server...")
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"params":  c.Query("1"),
		})
	})
	/* 	r.GET("/qE", query)
	   	r.GET("/insert", insertRandomId)
	   	r.GET("/getCount", getCount) */

	// ÖZGÜR BEY'İN YAZDIKLARI
	r.GET("/execute", executeParam)
	r.GET("/getCountsByParsing", getCountsByParsing)
	r.GET("/getRowsWithLimit", getRowsWithLimit)
	r.GET("/getMaxId", getMaxId)
	r.GET("/getRandom", getRandomRowById)

	go func() {
		if err := r.Run(":9000"); err != nil {
			fmt.Fprintf(os.Stderr, "Error starting server: %v\n", err)
		}
		fmt.Println("Server started")
	}()

	fmt.Println("Press Ctrl+C to stop the server")
	select {}
}

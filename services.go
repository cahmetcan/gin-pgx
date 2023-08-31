package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func qE(c *gin.Context) {
	// randomId := strconv.Itoa(rand.Intn(100))
	startQuery := time.Now()

	/* 	query := (fmt.Sprintf(`
		insert into test (id) values (%s) returning id;
	`, randomId)) */

	// get query from params
	query := c.Query("q")

	if query == "" {
		c.JSON(400, gin.H{
			"success": "false",
			"message": "No query provided",
		})
		return
	}
	// fmt.Print("Inserting ", randomId, " into test table... ")
	fmt.Println("Query:", query)

	pool, _, err := dbConnect()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		c.JSON(500, gin.H{
			"success": "false",
			"message": err.Error(),
		})
		return
	}

	result := eQuery(query, pool)

	if err != nil {
		fmt.Println("Error inserting into test table:", err)
		c.JSON(500, gin.H{
			"success": "false",
			"message": err.Error(),
		})
		return
	}

	fmt.Println("Time to query:", time.Since(startQuery))

	c.JSON(200, gin.H{
		"success": "true",
		"message": result,
	})
}

/*
blockchain'den rpc ile logları alıp datayı parse edip anlamlandırıp db'ye yazdırmak
websocket ile blockchain'den gelen datayı dinlemek
*/

func i(c *gin.Context) {
	randomId := strconv.Itoa(rand.Intn(100))
	startQuery := time.Now()
	fmt.Println("Inserting ", randomId, " into test table... ")

	query := (fmt.Sprintf(`
		insert into test (id) values (%s) returning id;
	`, randomId))

	pool, _, err := dbConnect()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		c.JSON(500, gin.H{
			"success": "false",
			"message": err.Error(),
		})
		return
	}

	result := eQuery(query, pool)

	if err != nil {
		fmt.Println("Error inserting into test table:", err)
		c.JSON(500, gin.H{
			"success": "false",
			"message": err.Error(),
		})
		return
	}

	fmt.Println("Time to query:", time.Since(startQuery))

	c.JSON(200, gin.H{
		"success": "true",
		"message": result,
	})

}

func s(c *gin.Context) {
	query := "SELECT * FROM test"
	_, conn, err := dbConnect()

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		c.JSON(500, gin.H{
			"success": "false",
			"message": err.Error(),
		})
		return
	}

	result := execQuery(query, conn)
	resultString := fmt.Sprintf("%v", result)

	if err != nil {
		fmt.Println("Error inserting into test table:", err)
		c.JSON(500, gin.H{
			"success": "false",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": "true",
		"message": resultString,
	})
}

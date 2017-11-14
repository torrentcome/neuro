package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	var r = gin.Default()
	v1 := r.Group("/api/v1/")
	{
		// c
		v1.PUT("/c/", insertOne)
		// r
		v1.GET("/r/", getAll)
		// u
		v1.POST("/u/:id", updateOne)
		// d
		v1.DELETE("/d/:id", deleteOne)
	}
	r.Run()
}

func getAll(c *gin.Context) {
	fmt.Print("getAll")
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func insertOne(c *gin.Context) {
	fmt.Print("insertOne")
}
func updateOne(c *gin.Context) {
	id := c.Param("id")
	fmt.Printf("updateOne %v", id)
}
func deleteOne(c *gin.Context) {
	id := c.Param("id")
	fmt.Printf("deleteOne %v", id)
}

package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	. "../model"
)

var (
	db  *gorm.DB
	err error
)

func main() {
	initDB()
	initWS()
}

func initWS() {
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

func initDB() {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&TicTac{})
}

func getAll(c *gin.Context) {
	var tictac []TicTac
	if err := db.Find(&tictac).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, tictac)
	}

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

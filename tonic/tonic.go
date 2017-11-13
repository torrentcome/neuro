package main

import "github.com/gin-gonic/gin"

func main() {
	var r = gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello hello World")
	})
	r.Run()
}

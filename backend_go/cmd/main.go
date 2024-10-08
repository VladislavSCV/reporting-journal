package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Logger middleware will write the logs to gin.DefaultWriter stream
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run(":8000")
	if err != nil {
		return
	}
}

package main

import (
	"http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		log.Debug("receive ping...")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	http.ListenAndServe(":8888", router())
}
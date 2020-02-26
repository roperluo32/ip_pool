package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
	"ip_proxy/producer"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		log.Println("receive ping...")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	producer.Add(3, 5)
	http.ListenAndServe(":8888", r)
}
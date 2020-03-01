package main

import (
	"github.com/gin-gonic/gin"
	"ip_proxy/component/log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		log.Info("receive ping...")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	http.ListenAndServe(":8888", r)
}

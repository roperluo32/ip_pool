package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ip_proxy/abstract/storage"
	"ip_proxy/component/config"
	"ip_proxy/component/log"
	"ip_proxy/component/storage/redisstorage"
	"net/http"
)

var _storage storage.ProxyStorage

func init() {
	config.Init("conf", ".")
	_storage = redisstorage.NewReidsSaver()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		log.Info("receive ping...")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/get", func(c *gin.Context) {
		domain := c.Query("domain")
		if domain == "" {
			c.String(200, "fail.not effect domain")
		}
		proxy, err := _storage.GetOneValidProxy(domain)
		if err != nil {
			c.String(200, "fail")
			return
		}
		log.Infof("get proxy req.domain:%v, proxy:%v", domain, proxy)
		c.String(200, fmt.Sprintf("%v:%v", proxy.IP, proxy.Port))
		return
	})

	http.ListenAndServe(":8888", r)
}

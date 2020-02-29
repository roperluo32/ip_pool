package main

import (
	"ip_proxy/config"
	"ip_proxy/getter"
	"ip_proxy/producer"
	"ip_proxy/redisstorage"
	"ip_proxy/validator"
	"log"
)

func main() {
	config.Init("conf", "..")
	redisSaver := redisstorage.NewReidsSaver()
	xdlGetter := getter.NewXunDaiLiGetter()

	prod := producer.NewProducer(redisSaver)
	prod.RegisterProxyGetter(xdlGetter)

	validate := validator.NewValidator(redisSaver, &validator.HTTPChecker{})

	go prod.Run()
	go validate.Run()

	log.Printf("pool main start\n")
	select {} //阻止main函数退出
}

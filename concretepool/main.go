package main

import (
	"ip_proxy/concretecmpt/checker/httpchecker"
	"ip_proxy/concretecmpt/config"
	"ip_proxy/concretecmpt/getter/xdaili"
	"ip_proxy/concretecmpt/storage/redisstorage"
	"ip_proxy/pool/producer"
	"ip_proxy/pool/validator"
	"log"
)

func main() {
	config.Init("conf", "..")
	redisSaver := redisstorage.NewReidsSaver()
	xdlGetter := xdaili.NewXunDaiLiGetter()

	prod := producer.NewProducer(redisSaver)
	prod.RegisterProxyGetter(xdlGetter)

	validate := validator.NewValidator(redisSaver, &httpchecker.HTTPChecker{})

	go prod.Run()
	go validate.Run()

	log.Printf("pool main start\n")
	select {} //阻止main函数退出
}

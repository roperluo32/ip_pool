package main

import (
	"ip_proxy/component/checker/httpchecker"
	"ip_proxy/component/config"
	"ip_proxy/component/getter/xdaili"
	"ip_proxy/component/log"
	// "ip_proxy/component/log/logrusger"
	"ip_proxy/component/storage/redisstorage"
	"ip_proxy/pool/producer"
	"ip_proxy/pool/validator"
)

func main() {
	config.Init("conf", "..")
	// 初始化日志
	// logIns := logrusger.NewLogrusLogger()
	// log.SetLogger(logIns)
	// 初始化底层部件
	redisSaver := redisstorage.NewReidsSaver()
	xdlGetter := xdaili.NewXunDaiLiGetter()
	// 初始化代理生产者
	prod := producer.NewProducer(redisSaver)
	prod.RegisterProxyGetter(xdlGetter)
	// 初始化代理验证器
	validate := validator.NewValidator(redisSaver, &httpchecker.HTTPChecker{})

	go prod.Run()
	go validate.Run()

	log.Infof("pool main start\n")
	select {} //阻止main函数退出
}

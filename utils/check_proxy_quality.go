package main

import (
	"fmt"
	"ip_proxy/component/checker/httpchecker"
	"ip_proxy/component/config"
	"ip_proxy/component/getter/xdaili"
	"ip_proxy/component/log"
	"ip_proxy/model"
	"sync"
	"time"
)

type Proxy struct {
	proxy model.IPItem
	start time.Time
	end   time.Time
	alive bool
}

func (p *Proxy) String() string {
	return fmt.Sprintf("proxy:%v:%v, duration:%v, alive:%v, start:%v", p.proxy.IP, p.proxy.Port, p.end.Sub(p.start), p.alive, p.start)
}

var checkDomain = "www.douban.com"

func getNumProxies(count int) []*Proxy {
	xdl := xdaili.NewXunDaiLiGetter()
	var proxies []*Proxy

	items, err := xdl.GetProxyIPs()
	if err != nil {
		log.Panicf("xdl get proxy ip fail.err:%v", err)
	}
	for _, item := range items {
		log.Printf("get xdaili raw proxy:%v", item)
		proxies = append(proxies, &Proxy{
			proxy: item,
			start: time.Now(),
			alive: true,
		})
	}

	return proxies
}

func checkOneLoop(proxies []*Proxy) {
	checker := &httpchecker.HTTPChecker{}
	var wg sync.WaitGroup
	for _, p := range proxies {
		if p.alive == false {
			continue
		}
		wg.Add(1)
		go func(p *Proxy) {
			defer wg.Done()
			isValid, err := checker.CheckProxyValid(checkDomain, p.proxy)
			if err != nil {
				log.Errorf("check proxy fail.proxy:%v, err:%v", p.proxy, err)
			}
			if isValid == false {
				p.end = time.Now()
				p.alive = false
			}
		}(p)
	}

	wg.Wait()
}

func run(proxies []*Proxy) {
	endTicker := time.NewTicker(10 * time.Minute)
	interTicker := time.NewTicker(8 * time.Second)
	defer endTicker.Stop()
	defer interTicker.Stop()

	for {
		select {
		case <-endTicker.C:
			return
		case <-interTicker.C:
			checkOneLoop(proxies)
		}
	}
}

func statistic(proxies []*Proxy) {
	totalNum := len(proxies)
	goodNum := 0
	var averageUseTime time.Duration
	var totalUseTime time.Duration
	now := time.Now()
	for _, p := range proxies {
		if p.alive {
			goodNum++
			totalUseTime += (now.Sub(p.start))
			continue
		} else {
			totalUseTime += p.end.Sub(p.start)
		}
	}
	averageUseTime = totalUseTime / time.Duration(totalNum)

	log.Printf("total num:%v.still good num:%v, bad num:%v", totalNum, goodNum, totalNum-goodNum)
	log.Printf("totalUseTime:%v, averageUseTime:%v", totalUseTime, averageUseTime)
	for _, p := range proxies {
		log.Printf("%v", p)
	}
}

func main() {
	config.Init("conf", "..")

	var proxies []*Proxy
	proxies = getNumProxies(2)

	run(proxies)

	statistic(proxies)
}

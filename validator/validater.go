package validator

import (
	// "github.com/asmcos/requests"
	"ip_proxy/config"
	"log"
	"sync"
	"time"
)

// Validator 代理校验器
type Validator struct {
	proxyStorage ProxyStorage
	proxyChecker ProxyChecker
	// quit
	quit    chan int
	domains []string
}

// Run 运行代理校验器
func (va *Validator) Run() {
	rawTicker := time.NewTicker(time.Duration(config.C.Validator.RawInterval) * time.Millisecond)
	validTicker := time.NewTicker(time.Duration(config.C.Validator.ValidInterval) * time.Millisecond)

	for {
		select {
		case <-va.quit:
			log.Println("receive quit signal. quit...")
			return
		case <-rawTicker.C:
			va.doRawCheck()
		case <-validTicker.C:
			va.doValidCheck()
		}
	}
}

func (va *Validator) doRawCheck() {
	var wg sync.WaitGroup
	for _, d := range va.domains {
		wg.Add(1)
		go func(domain string) {
			defer wg.Done()
			// 获取一个原始proxy
			proxy, err := va.proxyStorage.GetOneRawProxy(domain)
			if err != nil {
				log.Printf("[ERROR]get raw proxy for domain:%v fail.err:%+v\n", domain, err)
				return
			}
			if proxy.IP == "" {
				// log.Printf("[WARN] domain:%v don't have raw ip now\n", domain)
				return
			}

			// 检查是否有效
			isValid, err := va.proxyChecker.CheckProxyValid(domain, proxy)
			if err != nil {
				log.Printf("[ERROR]check raw proxy:%v for domain:%v fail.err:%+v\n", proxy, domain, err)
			}
			// 保存有效的代理
			if isValid {
				if err := va.proxyStorage.SaveValidProxy(domain, proxy); err != nil {
					log.Printf("[ERROR]save valid proxy fail.domain:%v, proxy:%v, err:%+v", domain, proxy, err)
					return
				}
			}

			// 从原始proxy池子中删掉
			if err := va.proxyStorage.DeleteRawProxy(domain, proxy, isValid); err != nil {
				log.Printf("DeleteRawProxy fail.err:%+v, domain:%v, proxy:%v\n", err, domain, proxy)
				return
			}
		}(d)

		wg.Wait()
	}
}

func (va *Validator) doValidCheck() {
	var wg sync.WaitGroup
	for _, d := range va.domains {
		wg.Add(1)
		go func(domain string) {
			defer wg.Done()
			// 获取一个valid proxy
			proxy, err := va.proxyStorage.GetOneValidProxy(domain)
			if err != nil {
				log.Printf("[ERROR]get valid proxy for domain:%v fail.err:%+v\n", domain, err)
				return
			}
			if proxy.IP == "" {
				// log.Printf("[WARN] domain:%v don't have valid ip now\n", domain)
				return
			}

			// 检查是否有效
			isValid, err := va.proxyChecker.CheckProxyValid(domain, proxy)
			if err != nil {
				log.Printf("[ERROR]check valid proxy:%v for domain:%v fail.err:%+v\n", proxy, domain, err)
			}
			// 无效的ip删除之
			if isValid == false {
				if err := va.proxyStorage.DeleteValidProxy(domain, proxy, isValid); err != nil {
					log.Printf("DeleteValidProxy fail.err:%+v, domain:%v, proxy:%v\n", err, domain, proxy)
					return
				}
			}
		}(d)

		wg.Wait()
	}
}

// Stop 停止
func (va *Validator) Stop() {
	va.quit <- 1
	return
}

var once sync.Once
var _validator Validator

// NewValidator 新建一个proxy校验器
func NewValidator(storage ProxyStorage, checker ProxyChecker) *Validator {
	once.Do(func() {
		_validator = Validator{
			proxyStorage: storage,
			proxyChecker: checker,
			quit:         make(chan int),
			domains:      config.C.Domains,
		}
	})

	return &_validator
}

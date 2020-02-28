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
	for {
		select {
		case <-va.quit:
			log.Println("receive quit signal. quit...")
			return
		default:
			break
		}
		va.oneLoop()
		time.Sleep(time.Duration(config.C.Validator.Interval) * time.Millisecond)
	}
}

func (va *Validator) oneLoop() {
	var wg sync.WaitGroup
	for _, d := range va.domains {
		wg.Add(1)
		go func(domain string) {
			defer wg.Done()
			// 获取一个原始proxy
			proxy, err := va.proxyStorage.GetOneRawProxy(domain)
			if err != nil {
				log.Printf("[WARN]get proxy for domain:%v fail.err:%+v\n", domain, err)
				return
			}

			// 检查是否有效
			isValid, err := va.proxyChecker.CheckRawProxy(domain, proxy)
			if err != nil {
				log.Printf("[ERROR]check proxy:%v for domain:%v fail.err:%+v\n", proxy, domain, err)
			}
			// 保存有效的代理
			if isValid {
				if err := va.proxyStorage.SaveValidProxy(domain, proxy); err != nil {
					log.Printf("[ERROR]save valid proxy fail.domain:%v, proxy:%v, err:%+v", domain, proxy, err)
					return
				}
			}

			// 从原始proxy池子中删掉
			va.proxyStorage.DeleteRawProxy(domain, proxy, isValid)
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

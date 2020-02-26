package producer

import (
	"time"
	"log"
)

// Producer 代理产生器
type Producer struct {
	// proxy getter的记录器，记录运行状态
	recorders []*getterRecorder		
	// 存储器
	saver ProxySaver
	// quit
	quit chan int
}

// 记录getter的运行状态
type getterRecorder struct {
	ProxyGetter
	// 上次运行的时间
	lastRunTime time.Time   
}

// GetNextRunTime 获取下一次运行的时间
func (gr *getterRecorder) GetNextRunTime() time.Time{
	return gr.lastRunTime.Add(gr.GetInterval())
}

// RegisterProxyGetter 注册proxy getter
func (p *Producer) RegisterProxyGetter(getter ProxyGetter) {
	p.recorders = append(p.recorders, &getterRecorder{ProxyGetter: getter, lastRunTime: time.Now()})
}

// Run 运行
func (p *Producer) Run(){
	// 100ms loop一次
	interval := 100 * time.Millisecond
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <- ticker.C:
			p.oneLoop()
		case <- p.quit:
			log.Println("receive quit signal. quit...")
			return
		}
	}
}

func (p *Producer) oneLoop() {
	now := time.Now()
	for _, recorder := range p.recorders {
		if now.Before(recorder.GetNextRunTime()) {
			continue
		}
		// 执行网络操作前赋值，避免重复执行
		recorder.lastRunTime = now	
		// 拉取代理ip
		go func() {
			ipItems, err := recorder.GetProxyIPs()
			if err != nil {
				log.Printf("[ERROR]: get proxy ip fail.err:%v, recorder:%v\n", err, recorder)
				return
			}
			if err := p.saver.SaveIpItems(ipItems); err != nil{
				log.Printf("[ERROR] save ip items fail.items:%v, err:%v\n", ipItems, err)
				return
			}
		}()
	}
}

// Stop
func (p *Producer) Stop() {
	p.quit <- 1
	return 
}

// NewProducer 新建一个proxy producer
func NewProducer(s ProxySaver) *Producer {
	return &Producer{
		saver: s,
		quit: make(chan int),
	}
}

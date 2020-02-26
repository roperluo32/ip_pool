package producer

import (
	"time"
)

// ProxyGetter 代理获取器
type ProxyGetter interface {
	// GetInterval 获取毫秒。代表多少毫秒可以执行一次
	GetInterval() time.Duration
	// GetProxyIPs 获取代理ip
	GetProxyIPs() ([]IPItem, error)
}

// Saver 存储器
type ProxySaver interface {
	// SaveIpItems保存proxy ip
	SaveIpItems([]IPItem) error
	// 获取存储中的proxy ip总数
	GetTotalNum() int
}
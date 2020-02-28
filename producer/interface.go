package producer

import (
	"ip_proxy/model"
	"time"
)

// ProxyGetter 代理获取器
type ProxyGetter interface {
	// GetInterval 获取毫秒。代表多少毫秒可以执行一次
	GetInterval() time.Duration
	// GetProxyIPs 获取代理ip
	GetProxyIPs() ([]model.IPItem, error)
}

// ProxySaver 存储器
type ProxySaver interface {
	// SaveIPItems保存proxy ip
	SaveIPItems([]model.IPItem) error
	// 获取存储中的raw proxy ip总数
	GetTotalNum() (int, error)
}

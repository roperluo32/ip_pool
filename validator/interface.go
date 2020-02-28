package validator

import (
	"ip_proxy/model"
)

// ProxyStorage 代理存取器，从存储storage中获取代理和修改，保存有效代理
type ProxyStorage interface {
	// 获取一个原始proxy ip
	GetOneRawProxy(domain string) (model.IPItem, error)
	// 删除一个原始proxy ip.isValid用来告诉Modifier，这个代理ip是否是有效的，以帮助它去决策是否真正从代理池中删除
	DeleteRawProxy(domain string, proxy model.IPItem, isValid bool) error
	// 保存一个有效proxy ip
	SaveValidProxy(domain string, proxy model.IPItem) error
}

// ProxyChecker  代理检查器，检查代理ip是否有效
type ProxyChecker interface {
	// 检查一个原始proxy ip是否有效
	CheckRawProxy(domain string, proxy model.IPItem) (bool, error)
}

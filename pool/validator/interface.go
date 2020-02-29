package validator

import (
	"ip_proxy/model"
)

// ProxyChecker  代理检查器，检查代理ip是否有效
type ProxyChecker interface {
	// 检查一个原始proxy ip是否有效
	CheckProxyValid(domain string, proxy model.IPItem) (bool, error)
}

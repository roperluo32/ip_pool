package storage

import (
	"ip_proxy/model"
)

// ProxySaver 存储器
type ProxySaver interface {
	// SaveIPItems保存proxy ip
	SaveIPItems([]model.IPItem) error
	// 获取存储中的raw proxy ip总数
	GetTotalNum() (int, error)
}

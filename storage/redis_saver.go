package storage

import (
	"ip_proxy/producer"
	"github.com/garyburd/redigo/redis"
)

type RedisSaver struct {
	
}

func (rs *RedisSaver) SaveIpItems(items []producer.IPItem) error {

	return nil
}

func (rs *RedisSaver) GetTotalNum() int {
	return 0
}
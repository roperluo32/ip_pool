package storage

import (
	"ip_proxy/producer"
	"github.com/garyburd/redigo/redis"
	"sync"
)

var once sync.Once

type RedisSaver struct {
	c redis.Conn
}

func (rs *RedisSaver) openRedisConn() error {

}
func (rs *RedisSaver) SaveIpItems(items []producer.IPItem) error {
	once.Do()
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	return nil
}

func (rs *RedisSaver) GetTotalNum() int {
	return 0
}
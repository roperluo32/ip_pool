package saver

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/pkg/errors"
	"ip_proxy/config"
	"ip_proxy/producer"
	"log"
	"sync"
)

var once sync.Once

// RedisSaverInstance redis连接实例
type RedisSaverInstance struct {
	conn     redis.Conn
	redisKey string
}

// SaveIPItems 保存代理
func (rs *RedisSaverInstance) SaveIPItems(items []producer.IPItem) error {
	for _, item := range items {
		if item.IP == "" || item.Port <= 0 {
			log.Printf("illegal IPItem:%v\n", item)
			continue
		}

		ipPort := fmt.Sprintf("%s:%d", item.IP, item.Port)
		if err := rs.conn.Send("lpush", rs.redisKey, ipPort); err != nil {
			log.Printf("save proxy to redis fail.proxy:%v, err:%v\n", item, err)
			continue
		}
	}
	return nil
}

// GetTotalNum 返回总的代理数
func (rs *RedisSaverInstance) GetTotalNum() (int, error) {
	totalNum, err := redis.Int(rs.conn.Do("llen", rs.redisKey))
	if err != nil {
		return -1, errors.Wrap(err, "get redis proxy total num fail")
	}

	return totalNum, err
}

// Clear 删除所有代理，用于test
func (rs *RedisSaverInstance) Clear() error {
	_, err := rs.conn.Do("del", rs.redisKey)
	if err != nil {
		return errors.Wrap(err, "clear redis proxy fail")
	}

	return nil
}

// SetRedisKey 设置存储proxy的redis key
func (rs *RedisSaverInstance) SetRedisKey(key string) error {
	rs.redisKey = key

	return nil
}

var _saverInstance RedisSaverInstance

// NewReidsSaver 新建单例实例
func NewReidsSaver() producer.ProxySaver {
	once.Do(func() {
		conn, err := redis.Dial("tcp", config.C.Redis.URL)
		if err != nil {
			log.Panicf("redis conn fail.err:%v\n", err)
		}

		if config.C.Redis.Password != "" {
			err = conn.Send("auth", config.C.Redis.Password)
			if err != nil {
				log.Panicf("redis auth fail.err:%v\n", err)
			}
		}

		_saverInstance.conn = conn
		_saverInstance.redisKey = "raw_proxies"
	})

	return &_saverInstance
}

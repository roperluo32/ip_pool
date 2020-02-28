package redisstorage

import (
	"github.com/garyburd/redigo/redis"
	"github.com/pkg/errors"
	"ip_proxy/config"
	"ip_proxy/producer"
	"log"
	"sync"
)

var once sync.Once

// RedisStorage redis连接实例
type RedisStorage struct {
	conn    redis.Conn
	domains []string //支持按域名存储多份代理IP
}

// Clear 删除所有代理，用于test
func (rs *RedisStorage) Clear() error {
	var err error
	for _, domain := range rs.domains {
		if _, err := rs.conn.Do("del", domain); err != nil {
			err = errors.Wrapf(err, "clear redis proxy for domain:%v fail", domain)
			continue
		}
	}

	return err
}

var _saverInstance RedisStorage

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
		_saverInstance.domains = config.C.Domains
	})

	return &_saverInstance
}

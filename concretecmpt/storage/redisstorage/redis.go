package redisstorage

import (
	"github.com/garyburd/redigo/redis"
	"github.com/pkg/errors"
	"ip_proxy/concretecmpt/config"
	"log"
	"sync"
)

// RedisStorage redis连接实例
type RedisStorage struct {
	conn     redis.Conn
	domains  []string //支持按域名存储多份代理IP
	rawKey   string   // 存放未检验代理的redis key
	validKey string   // 存放已检验代理的redis key
}

// Clear 删除所有代理，用于test
func (rs *RedisStorage) Clear() error {
	var err error
	for _, domain := range rs.domains {
		if _, err := rs.conn.Do("del", domain); err != nil {
			err = errors.Wrapf(err, "clear redis raw proxy for domain:%v fail", domain)
			continue
		}
		if _, err := rs.conn.Do("del", validProxyKey(domain)); err != nil {
			err = errors.Wrapf(err, "clear redis valid proxy for domain:%v fail", domain)
			continue
		}
	}

	return err
}

var once sync.Once
var _saverInstance RedisStorage

// NewReidsSaver 新建单例实例
func NewReidsSaver() *RedisStorage {
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
		_saverInstance = RedisStorage{
			conn:     conn,
			domains:  config.C.Domains,
			rawKey:   "raw_proxies",
			validKey: "valid_proxies",
		}
	})

	return &_saverInstance
}

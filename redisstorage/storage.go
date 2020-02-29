package redisstorage

import (
	"github.com/garyburd/redigo/redis"
	"github.com/pkg/errors"
	"ip_proxy/config"
	"ip_proxy/model"
	"log"
	"math/rand"
)

// GetOneRawProxy 获取一个原始proxy ip
func (rs *RedisStorage) GetOneRawProxy(domain string) (model.IPItem, error) {
	items, err := redis.Values(rs.conn.Do("hkeys", domain))
	if err != nil {
		return model.IPItem{}, errors.Wrap(err, "redis hkeys fail")
	}
	if len(items) == 0 {
		return model.IPItem{}, nil
	}

	proxy := items[rand.Intn(len(items))]
	key := string(proxy.([]byte))
	return stringToIPItem(key), nil
}

func (rs *RedisStorage) shouldDeleteRawProxy(domain string, key string, isValid bool) (bool, error) {
	if isValid { //代理ip是有效的，那么就直接从raw proxy中删除
		return true, nil
	}

	// 代理ip无效，并且重试超过N次（可配置），也该删除了
	tryTimes, err := redis.Int(rs.conn.Do("hget", domain, key))
	log.Printf("trytimes:%v, domain:%v, key:%v\n", tryTimes, domain, key)
	if err != nil {
		return false, errors.Wrapf(err, "hget fail.domain:%v, key:%v", domain, key)
	}
	if tryTimes+1 >= config.C.Redis.MaxTryTimes {
		return true, nil
	}

	return false, nil
}

// DeleteRawProxy 删除一个原始proxy ip.isValid用来告诉Modifier，这个代理ip是否是有效的，以帮助它去决策是否真正从代理池中删除
func (rs *RedisStorage) DeleteRawProxy(domain string, proxy model.IPItem, isValid bool) error {
	key := ipItemToString(proxy)
	canDelete, err := rs.shouldDeleteRawProxy(domain, key, isValid)
	if err != nil {
		return errors.Wrap(err, "shouldDeleteRawProxy return fail")
	}
	log.Printf("DeleteRawProxy.domain:%v, proxy:%v, canDelete:%v, isValid:%v\n", domain, proxy, canDelete, isValid)
	if canDelete {
		log.Printf("delete key:%v, domain:%v\n", key, domain)
		if _, err := rs.conn.Do("hdel", domain, key); err != nil {
			return errors.Wrap(err, "hdel fail")
		}
		return nil
	}
	if _, err := rs.conn.Do("hincrby", domain, key, 1); err != nil {
		return errors.Wrap(err, "hincrby fail")
	}

	return nil
}

// SaveValidProxy 保存一个有效proxy ip
func (rs *RedisStorage) SaveValidProxy(domain string, proxy model.IPItem) error {
	ipPort := ipItemToString(proxy)

	if err := rs.conn.Send("hmset", validProxyKey(domain), ipPort, 0); err != nil {
		return errors.Wrap(err, "hmset fail")
	}

	return nil
}

// GetNumOfValid 获取有效代理的数量
func (rs *RedisStorage) GetNumOfValid(domain string) (int, error) {
	totalNum, err := redis.Int(rs.conn.Do("hlen", validProxyKey(domain)))
	if err != nil {
		return -1, errors.Wrapf(err, "get redis valid proxy total num for domain:%v fail", domain)
	}

	return totalNum, err
}

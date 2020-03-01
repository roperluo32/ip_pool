package redisstorage

import (
	"github.com/garyburd/redigo/redis"
	"github.com/pkg/errors"
	"ip_proxy/component/config"
	"ip_proxy/component/log"
	"ip_proxy/model"
	"math/rand"
)

// GetOneRawProxy 获取一个原始proxy ip
func (rs *RedisStorage) GetOneRawProxy(domain string) (model.IPItem, error) {
	proxy, err := rs.getRandFieldFromHashMap(domain)
	if err != nil {
		return model.IPItem{}, errors.Wrap(err, "getRandFieldFromHashMap fail")
	}
	if proxy == "" {
		return model.IPItem{}, nil
	}

	return stringToIPItem(proxy), nil
}

func (rs *RedisStorage) shouldDelete(hmkey string, key string, isValid bool) (bool, error) {
	if isValid { //代理ip是有效的，那么就直接从raw proxy中删除
		log.Infof("proxy is valid.should delete from raw.domain:%v, key:%v", hmkey, key)
		return true, nil
	}

	// 代理ip无效，并且重试超过N次（可配置），也该删除了
	tryTimes, err := redis.Int(rs.conn.Do("hget", hmkey, key))
	// log.Infof("trytimes:%v, domain:%v, key:%v\n", tryTimes, hmkey, key)
	if err != nil {
		return false, errors.Wrapf(err, "hget fail.domain:%v, key:%v", hmkey, key)
	}
	if tryTimes+1 >= config.C.Redis.MaxTryTimes {
		log.Infof("proxy tries max times.should delete from raw.domain:%v, key:%v", hmkey, key)
		return true, nil
	}

	return false, nil
}

func (rs *RedisStorage) deleteOneProxy(hmKey string, key string, isValid bool) error {
	canDelete, err := rs.shouldDelete(hmKey, key, isValid)
	if err != nil {
		return errors.Wrap(err, "shouldDelete return fail")
	}

	if canDelete {
		if _, err := rs.conn.Do("hdel", hmKey, key); err != nil {
			return errors.Wrap(err, "hdel fail")
		}
		return nil
	}
	if _, err := rs.conn.Do("hincrby", hmKey, key, 1); err != nil {
		return errors.Wrap(err, "hincrby fail")
	}

	return nil
}

// DeleteRawProxy 删除一个原始proxy ip.isValid用来告诉Modifier，这个代理ip是否是有效的，以帮助它去决策是否真正从代理池中删除
func (rs *RedisStorage) DeleteRawProxy(domain string, proxy model.IPItem, isValid bool) error {
	key := ipItemToString(proxy)
	return rs.deleteOneProxy(domain, key, isValid)
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

func (rs *RedisStorage) getRandFieldFromHashMap(hmKey string) (string, error) {
	items, err := redis.Values(rs.conn.Do("hkeys", hmKey))
	if err != nil {
		return "", errors.Wrap(err, "redis hkeys fail")
	}
	if len(items) == 0 {
		return "", nil
	}

	proxy := items[rand.Intn(len(items))]
	return string(proxy.([]byte)), nil
}

// GetOneValidProxy 获取一个有效proxy ip
func (rs *RedisStorage) GetOneValidProxy(domain string) (model.IPItem, error) {
	proxy, err := rs.getRandFieldFromHashMap(validProxyKey(domain))
	if err != nil {
		return model.IPItem{}, errors.Wrap(err, "getRandFieldFromHashMap fail")
	}
	if proxy == "" {
		return model.IPItem{}, nil
	}

	return stringToIPItem(proxy), nil
}

// DeleteValidProxy 删除一个有效proxy ip
func (rs *RedisStorage) DeleteValidProxy(domain string, proxy model.IPItem, isValid bool) error {
	key := ipItemToString(proxy)
	hmKey := validProxyKey(domain)
	return rs.deleteOneProxy(hmKey, key, isValid)
}

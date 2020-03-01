package redisstorage

import (
	"github.com/garyburd/redigo/redis"
	"github.com/pkg/errors"
	"ip_proxy/component/log"
	"ip_proxy/model"
)

// SaveIPItems 保存代理
func (rs *RedisStorage) SaveIPItems(items []model.IPItem) error {
	for _, item := range items {
		if item.IP == "" || item.Port <= 0 {
			log.Errorf("illegal IPItem:%v\n", item)
			continue
		}

		ipPort := ipItemToString(item)
		for _, domain := range rs.domains {
			if err := rs.conn.Send("hmset", domain, ipPort, 0); err != nil {
				log.Errorf("save proxy to redis domain:%v fail.proxy:%v, err:%v\n", domain, item, err)
				continue
			}
		}

	}
	return nil
}

// GetTotalNum 返回总的原始代理数
func (rs *RedisStorage) GetTotalNum() (int, error) {
	// 默认返回domains[0]的代理数量
	totalNum, err := redis.Int(rs.conn.Do("hlen", rs.domains[0]))
	if err != nil {
		return -1, errors.Wrapf(err, "get redis raw proxy total num for domain:%v fail", rs.domains[0])
	}

	return totalNum, err
}

// GetNumOfRaw 返回指定域名的原始代理数
func (rs *RedisStorage) GetNumOfRaw(domain string) (int, error) {
	totalNum, err := redis.Int(rs.conn.Do("hlen", domain))
	if err != nil {
		return -1, errors.Wrapf(err, "get redis raw proxy total num for domain:%v fail", domain)
	}

	return totalNum, err
}

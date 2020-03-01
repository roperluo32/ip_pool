package storage

import (
	"fmt"
	"ip_proxy/model"
	"math/rand"
	"time"
)

type MockStorage struct {
	proxies map[string][]model.IPItem
	valids  map[string][]model.IPItem
}

func (ml *MockStorage) Init() {
	ml.proxies = map[string][]model.IPItem{
		"www.ropertest.com": []model.IPItem{
			model.IPItem{IP: "127.0.0.1", Port: 8008, Typ: model.TypGaoNi},
			model.IPItem{IP: "192.168.0.2", Port: 18668, Typ: model.TypGaoNi},
			model.IPItem{IP: "17.291.192.6", Port: 6115, Typ: model.TypGaoNi},
			model.IPItem{IP: "39.27.92.38", Port: 2893, Typ: model.TypGaoNi},
		},
		"www.aaaa.cn": []model.IPItem{
			model.IPItem{IP: "58.19.0.23", Port: 37082, Typ: model.TypGaoNi},
		},
	}
	ml.valids = map[string][]model.IPItem{}
}

func (ml *MockStorage) GetOneRawProxy(domain string) (model.IPItem, error) {
	domain_proxies, ok := ml.proxies[domain]
	if ok == false {
		return model.IPItem{}, fmt.Errorf("domain:%v not supported", domain)
	}
	if len(domain_proxies) == 0 {
		return model.IPItem{}, nil
	}
	rand.Seed(time.Now().Unix())
	return domain_proxies[rand.Intn(len(domain_proxies))], nil
}

// 删除一个原始proxy ip
func (ml *MockStorage) DeleteRawProxy(domain string, proxy model.IPItem, valid bool) error {
	domain_proxies, ok := ml.proxies[domain]
	if ok == false {
		return fmt.Errorf("domain:%v not supported", domain)
	}
	for i, item := range domain_proxies {
		if item.IP == proxy.IP && item.Port == proxy.Port {
			ml.proxies[domain] = append(ml.proxies[domain][:i], ml.proxies[domain][i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("proxy:%v not exist", proxy)
}

func (ml *MockStorage) GetNumOfRaw(domain string) (int, error) {
	return len(ml.proxies[domain]), nil
}

// 保存一个有效proxy ip
func (ml *MockStorage) SaveValidProxy(domain string, proxy model.IPItem) error {
	_, ok := ml.valids[domain]
	if ok == false {
		ml.valids[domain] = []model.IPItem{}
	}

	ml.valids[domain] = append(ml.valids[domain], proxy)
	return nil
}

// GetOneValidProxy 获取一个有效proxy ip
func (ml *MockStorage) GetOneValidProxy(domain string) (model.IPItem, error) {
	domain_proxies, ok := ml.valids[domain]
	if ok == false {
		return model.IPItem{}, fmt.Errorf("domain:%v not exist", domain)
	}
	if len(domain_proxies) == 0 {
		return model.IPItem{}, nil
	}
	rand.Seed(time.Now().Unix())
	return domain_proxies[rand.Intn(len(domain_proxies))], nil
}

// 删除一个有效proxy ip
func (ml *MockStorage) DeleteValidProxy(domain string, proxy model.IPItem, isValid bool) error {
	domain_proxies, ok := ml.valids[domain]
	if ok == false {
		return fmt.Errorf("domain:%v not supported", domain)
	}
	for i, item := range domain_proxies {
		if item.IP == proxy.IP && item.Port == proxy.Port {
			ml.valids[domain] = append(ml.valids[domain][:i], ml.valids[domain][i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("proxy:%v not exist", proxy)
}

func (ml *MockStorage) GetNumOfValid(domain string) (int, error) {
	return len(ml.valids[domain]), nil
}

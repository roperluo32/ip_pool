package producer

import (
	"ip_proxy/model"
	"math/rand"
	"time"
)

type MockGetter struct {
}

var itemPool = []model.IPItem{
	model.IPItem{IP: "127.0.0.1", Port: 8008, Typ: model.TypGaoNi},
	model.IPItem{IP: "192.168.0.2", Port: 18668, Typ: model.TypGaoNi},
	model.IPItem{IP: "17.291.192.6", Port: 6115, Typ: model.TypGaoNi},
	model.IPItem{IP: "39.27.92.38", Port: 2893, Typ: model.TypGaoNi},
}

func (mg *MockGetter) GetInterval() time.Duration {
	return 800 * time.Millisecond
}

func (mg *MockGetter) GetProxyIPs() ([]model.IPItem, error) {
	rand.Seed(time.Now().Unix())
	return []model.IPItem{itemPool[rand.Intn(len(itemPool))]}, nil
}

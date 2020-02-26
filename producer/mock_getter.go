package producer

import (
	"time"
	"math/rand"
)

type MockGetter struct {

}

var itemPool = []IPItem {
	IPItem{ip:"127.0.0.1", port:8008, typ:TypGaoNi,},
	IPItem{ip:"192.168.0.2", port:18668, typ:TypGaoNi,},
	IPItem{ip:"17.291.192.6", port:6115, typ:TypGaoNi,},
	IPItem{ip:"39.27.92.38", port:2893, typ:TypGaoNi,},
}

func (mg *MockGetter) GetInterval() time.Duration{
	return 800 * time.Millisecond
}

func (mg *MockGetter) GetProxyIPs() ([]IPItem, error) {
	rand.Seed(time.Now().Unix())
	return []IPItem{itemPool[rand.Intn(len(itemPool))], }, nil
}
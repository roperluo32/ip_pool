package producer

import (
	"math/rand"
	"time"
)

type MockGetter struct {
}

var itemPool = []IPItem{
	IPItem{IP: "127.0.0.1", Port: 8008, Typ: TypGaoNi},
	IPItem{IP: "192.168.0.2", Port: 18668, Typ: TypGaoNi},
	IPItem{IP: "17.291.192.6", Port: 6115, Typ: TypGaoNi},
	IPItem{IP: "39.27.92.38", Port: 2893, Typ: TypGaoNi},
}

func (mg *MockGetter) GetInterval() time.Duration {
	return 800 * time.Millisecond
}

func (mg *MockGetter) GetProxyIPs() ([]IPItem, error) {
	rand.Seed(time.Now().Unix())
	return []IPItem{itemPool[rand.Intn(len(itemPool))]}, nil
}

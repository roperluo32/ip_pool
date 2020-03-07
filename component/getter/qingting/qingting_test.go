package qingting

import (
	"ip_proxy/component/log"
	"testing"
	"time"

	"ip_proxy/component/config"

	"github.com/stretchr/testify/assert"
)

func TestQingTingBasic(t *testing.T) {
	config.Init("test.conf", "../../..")

	qtGetter := NewQingTingGetter()
	log.Debugf("qtGetter:%v\n", qtGetter)
	interval := qtGetter.GetInterval()
	assert.Equal(t, 10*time.Second, interval)

	items, err := qtGetter.GetProxyIPs()
	log.Debugf("items:%v\n", items)
	assert.Nil(t, err)
}

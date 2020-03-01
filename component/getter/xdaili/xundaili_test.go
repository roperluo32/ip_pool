package xdaili

import (
	"ip_proxy/component/log"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"ip_proxy/component/config"
)

func TestXunDaiLiBasic(t *testing.T) {
	config.Init("test.conf", "../../..")

	xdl := NewXunDaiLiGetter()
	log.Debugf("xdl:%v\n", xdl)
	interval := xdl.GetInterval()
	assert.Equal(t, 10*time.Second, interval)

	count, _ := strconv.Atoi(xdl.GetCount())
	items, err := xdl.GetProxyIPs()
	log.Debugf("items:%v\n", items)
	assert.Nil(t, err)
	assert.Equal(t, count, len(items))
}

package xdaili

import (
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"ip_proxy/concretecmpt/config"
)

func TestXunDaiLiBasic(t *testing.T) {
	config.Init("test.conf", "../../..")

	xdl := NewXunDaiLiGetter()
	log.Printf("xdl:%v\n", xdl)
	interval := xdl.GetInterval()
	assert.Equal(t, 10*time.Second, interval)

	count, _ := strconv.Atoi(xdl.GetCount())
	items, err := xdl.GetProxyIPs()
	log.Printf("items:%v\n", items)
	assert.Nil(t, err)
	assert.Equal(t, count, len(items))
}

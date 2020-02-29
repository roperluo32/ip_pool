package getter

import (
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"ip_proxy/config"
)

func TestXunDaiLiBasic(t *testing.T) {
	config.Init("test.conf", "..")

	getter := NewXunDaiLiGetter()
	xdl := getter.(*XunDaiLi)
	log.Printf("xdl:%v\n", xdl)
	interval := xdl.GetInterval()
	assert.Equal(t, 10*time.Second, interval)

	count, _ := strconv.Atoi(xdl.GetCount())
	items, err := xdl.GetProxyIPs()
	log.Printf("items:%v\n", items)
	assert.Nil(t, err)
	assert.Equal(t, count, len(items))
}

package redisstorage

import (
	"ip_proxy/config"
	"ip_proxy/model"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedisSaverBasic(t *testing.T) {
	config.Init("test.conf", "..")
	rs := NewReidsSaver()

	rs.SaveIPItems([]model.IPItem{
		model.IPItem{IP: "127.0.0.1", Port: 11118},
		model.IPItem{IP: "192.168.0.11", Port: 28808},
		model.IPItem{IP: "26.0.1.1", Port: 9278},
	})

	num, err := rs.GetTotalNum()
	log.Printf("num:%v\n", num)
	assert.Nil(t, err)
	assert.Equal(t, 3, num)

	err = rs.Clear()
	assert.Nil(t, err)
}

func TestRedisSaverZero(t *testing.T) {
	config.Init("test.conf", "..")
	rs := NewReidsSaver()
	// 保存一个0个代理
	rs.SaveIPItems([]model.IPItem{})

	num, err := rs.GetTotalNum()
	log.Printf("num:%v\n", num)
	assert.Nil(t, err)
	assert.Equal(t, 0, num)

	err = rs.Clear()
	assert.Nil(t, err)
}

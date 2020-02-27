package saver

import (
	"ip_proxy/config"
	"ip_proxy/producer"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedisSaverBasic(t *testing.T) {
	config.Init("conf", "..")
	saver := NewReidsSaver()
	rs := saver.(*RedisSaverInstance)
	rs.SetRedisKey("test_2")
	rs.SaveIPItems([]producer.IPItem{
		producer.IPItem{IP: "127.0.0.1", Port: 11118},
		producer.IPItem{IP: "192.168.0.11", Port: 28808},
		producer.IPItem{IP: "26.0.1.1", Port: 9278},
	})

	num, err := rs.GetTotalNum()
	log.Printf("num:%v\n", num)
	assert.Nil(t, err)
	assert.Equal(t, 3, num)

	err = rs.Clear()
	assert.Nil(t, err)
}

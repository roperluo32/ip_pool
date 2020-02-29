package redisstorage

import (
	"ip_proxy/concretecmpt/config"
	"ip_proxy/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedisValidatorBasic(t *testing.T) {
	config.Init("test.conf", "../../..")
	rs := NewReidsSaver()

	rs.SaveIPItems([]model.IPItem{
		model.IPItem{IP: "127.0.0.1", Port: 11118},
		model.IPItem{IP: "192.168.0.11", Port: 28808},
		model.IPItem{IP: "26.0.1.1", Port: 9278},
	})

	num, err := rs.GetTotalNum()
	assert.Nil(t, err)
	assert.Equal(t, 3, num)

	domain := "www.ropertest.com"
	item, err := rs.GetOneRawProxy(domain)
	assert.Nil(t, err)

	err = rs.SaveValidProxy(domain, item)
	assert.Nil(t, err)
	num, err = rs.GetNumOfValid(domain)
	assert.Nil(t, err)
	assert.Equal(t, 1, num)

	err = rs.DeleteRawProxy(domain, item, false)
	assert.Nil(t, err)
	num, err = rs.GetNumOfRaw(domain)
	assert.Nil(t, err)
	assert.Equal(t, 3, num)

	err = rs.DeleteRawProxy(domain, item, true)
	assert.Nil(t, err)
	num, err = rs.GetNumOfRaw(domain)
	assert.Nil(t, err)
	assert.Equal(t, 2, num)

	err = rs.Clear()
	assert.Nil(t, err)
}

func TestRedisValidatorEmpty(t *testing.T) {
	config.Init("test.conf", "../../..")
	rs := NewReidsSaver()

	rs.SaveIPItems([]model.IPItem{})

	num, err := rs.GetTotalNum()
	assert.Nil(t, err)
	assert.Equal(t, 0, num)

	domain := "www.ropertest.com"
	item, err := rs.GetOneRawProxy(domain)
	assert.Nil(t, err)
	assert.Equal(t, "", item.IP)

	err = rs.DeleteRawProxy(domain, item, false)
	assert.NotNil(t, err)
	num, err = rs.GetNumOfRaw(domain)
	assert.Nil(t, err)
	assert.Equal(t, 0, num)

	num, err = rs.GetNumOfValid(domain)
	assert.Nil(t, err)
	assert.Equal(t, 0, num)

	err = rs.Clear()
	assert.Nil(t, err)
}

func TestRedisValidatorDeleteRaw(t *testing.T) {
	config.Init("test.conf", "../../..")
	rs := NewReidsSaver()

	rs.SaveIPItems([]model.IPItem{
		model.IPItem{IP: "127.0.0.1", Port: 11118},
		model.IPItem{IP: "192.168.0.11", Port: 28808},
		model.IPItem{IP: "26.0.1.1", Port: 9278},
	})

	domain := "www.ropertest.com"
	item, err := rs.GetOneRawProxy(domain)
	assert.Nil(t, err)

	err = rs.DeleteRawProxy(domain, item, false)
	assert.Nil(t, err)
	num, err := rs.GetNumOfRaw(domain)
	assert.Nil(t, err)
	assert.Equal(t, 3, num)

	err = rs.DeleteRawProxy(domain, item, false)
	assert.Nil(t, err)
	num, err = rs.GetNumOfRaw(domain)
	assert.Nil(t, err)
	assert.Equal(t, 3, num)

	// 第三次删除（次数可配置），会真正的删掉
	err = rs.DeleteRawProxy(domain, item, false)
	assert.Nil(t, err)
	num, err = rs.GetNumOfRaw(domain)
	assert.Nil(t, err)
	assert.Equal(t, 2, num)

	// 获取一个新的item，执行一次valid=True的删除
	item, err = rs.GetOneRawProxy(domain)
	assert.Nil(t, err)
	err = rs.DeleteRawProxy(domain, item, true)
	assert.Nil(t, err)
	num, err = rs.GetNumOfRaw(domain)
	assert.Nil(t, err)
	assert.Equal(t, 1, num)

	err = rs.Clear()
	assert.Nil(t, err)
}

func TestRedisValidOperation(t *testing.T) {
	config.Init("test.conf", "../../..")
	rs := NewReidsSaver()

	rs.SaveIPItems([]model.IPItem{
		model.IPItem{IP: "127.0.0.1", Port: 11118},
		model.IPItem{IP: "192.168.0.11", Port: 28808},
		model.IPItem{IP: "26.0.1.1", Port: 9278},
	})

	domain := "www.ropertest.com"
	item, err := rs.GetOneRawProxy(domain)
	assert.Nil(t, err)

	// 保存一个valid proxy
	err = rs.SaveValidProxy(domain, item)
	assert.Nil(t, err)
	num, err := rs.GetNumOfValid(domain)
	assert.Nil(t, err)
	assert.Equal(t, 1, num)

	err = rs.DeleteRawProxy(domain, item, true)
	assert.Nil(t, err)
	num, err = rs.GetNumOfRaw(domain)
	assert.Nil(t, err)
	assert.Equal(t, 2, num)

	// 再保存一个valid proxy
	item, err = rs.GetOneRawProxy(domain)
	assert.Nil(t, err)

	err = rs.SaveValidProxy(domain, item)
	assert.Nil(t, err)
	num, err = rs.GetNumOfValid(domain)
	assert.Nil(t, err)
	assert.Equal(t, 2, num)

	err = rs.DeleteRawProxy(domain, item, true)
	assert.Nil(t, err)
	num, err = rs.GetNumOfRaw(domain)
	assert.Nil(t, err)
	assert.Equal(t, 1, num)

	// 获取一个valid proxy
	validProxy, err := rs.GetOneValidProxy(domain)
	assert.Nil(t, err)
	assert.NotEqual(t, "", validProxy.IP)
	// 检查valid的个数
	num, err = rs.GetNumOfValid(domain)
	assert.Nil(t, err)
	assert.Equal(t, 2, num)

	// 删除一个valid proxy
	err = rs.DeleteValidProxy(domain, validProxy, true)
	assert.Nil(t, err)
	num, err = rs.GetNumOfValid(domain)
	assert.Nil(t, err)
	assert.Equal(t, 1, num)

	err = rs.Clear()
	assert.Nil(t, err)
}

func TestRedisDeleteValid(t *testing.T) {
	config.Init("test.conf", "../../..")
	rs := NewReidsSaver()

	rs.SaveIPItems([]model.IPItem{
		model.IPItem{IP: "127.0.0.1", Port: 11118},
		model.IPItem{IP: "192.168.0.11", Port: 28808},
		model.IPItem{IP: "26.0.1.1", Port: 9278},
	})

	domain := "www.ropertest.com"
	item, err := rs.GetOneRawProxy(domain)
	assert.Nil(t, err)

	// 保存一个valid proxy
	err = rs.SaveValidProxy(domain, item)
	assert.Nil(t, err)
	num, err := rs.GetNumOfValid(domain)
	assert.Nil(t, err)
	assert.Equal(t, 1, num)

	err = rs.DeleteRawProxy(domain, item, true)
	assert.Nil(t, err)
	num, err = rs.GetNumOfRaw(domain)
	assert.Nil(t, err)
	assert.Equal(t, 2, num)

	// 再保存一个valid proxy
	item, err = rs.GetOneRawProxy(domain)
	assert.Nil(t, err)

	err = rs.SaveValidProxy(domain, item)
	assert.Nil(t, err)
	num, err = rs.GetNumOfValid(domain)
	assert.Nil(t, err)
	assert.Equal(t, 2, num)

	// 删除valid proxy（此时不会真的删除）
	err = rs.DeleteValidProxy(domain, item, false)
	assert.Nil(t, err)
	num, err = rs.GetNumOfValid(domain)
	assert.Nil(t, err)
	assert.Equal(t, 2, num)

	err = rs.DeleteValidProxy(domain, item, false)
	assert.Nil(t, err)
	num, err = rs.GetNumOfValid(domain)
	assert.Nil(t, err)
	assert.Equal(t, 2, num)
	//到第三次（可配置）会真正删除
	err = rs.DeleteValidProxy(domain, item, false)
	assert.Nil(t, err)
	num, err = rs.GetNumOfValid(domain)
	assert.Nil(t, err)
	assert.Equal(t, 1, num)

}

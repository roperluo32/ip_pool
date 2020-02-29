package validator

import (
	"github.com/stretchr/testify/assert"
	"ip_proxy/config"
	"testing"
	"time"
)

func TestValidatorBasic(t *testing.T) {
	config.Init("test.conf", "..")
	storage := &MockStorage{}
	storage.Init()

	checker := &MockChecker{}
	validor := NewValidator(storage, checker)
	go validor.Run()

	time.Sleep(2 * time.Second)

	validor.Stop()

	domain := "www.aaaa.cn"
	num, err := storage.GetNumOfRaw(domain)
	assert.Nil(t, err)
	assert.Equal(t, 0, num)
	num, err = storage.GetNumOfValid(domain)
	assert.Nil(t, err)
	assert.Equal(t, 0, num)
	proxy, err := storage.GetOneValidProxy(domain)
	assert.NotNil(t, err)
	assert.Equal(t, "", proxy.IP)

	domain = "www.ropertest.com"
	num, err = storage.GetNumOfRaw(domain)
	assert.Nil(t, err)
	assert.Equal(t, 0, num)
	num, err = storage.GetNumOfValid(domain)
	assert.Nil(t, err)
	assert.Equal(t, 2, num)
	proxy, err = storage.GetOneValidProxy(domain)
	assert.Nil(t, err)
	assert.NotEqual(t, "", proxy.IP)
	err = storage.DeleteValidProxy(domain, proxy, true)
	assert.Nil(t, err)
	num, err = storage.GetNumOfValid(domain)
	assert.Nil(t, err)
	assert.Equal(t, 1, num)
}

package validator

import (
	"github.com/stretchr/testify/assert"
	"ip_proxy/config"
	"testing"
	"time"
)

func TestValidatorBasic(t *testing.T) {
	config.Init("conf", "..")
	storage := &MockStorage{}
	storage.Init()

	checker := &MockChecker{}
	validor := NewValidator(storage, checker)
	go validor.Run()

	time.Sleep(2 * time.Second)

	validor.Stop()

	num, err := storage.GetNumOfDomain("www.zhihu.com")
	assert.Nil(t, err)
	assert.Equal(t, 0, num)
	num, err = storage.GetNumOfValid("www.zhihu.com")
	assert.Nil(t, err)
	assert.Equal(t, 0, num)

	num, err = storage.GetNumOfDomain("www.douban.com")
	assert.Nil(t, err)
	assert.Equal(t, 0, num)
	num, err = storage.GetNumOfValid("www.douban.com")
	assert.Nil(t, err)
	assert.Equal(t, 2, num)
}

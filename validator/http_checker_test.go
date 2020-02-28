package validator

import (
	"ip_proxy/config"
	"ip_proxy/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpCheckerBasic(t *testing.T) {
	config.Init("conf", "..")
	c := HTTPChecker{}
	isValid, err := c.CheckRawProxy("www.douban.com", model.IPItem{IP: "180.125.70.156", Port: 32572})
	assert.Nil(t, err)
	assert.Equal(t, isValid, true)
}

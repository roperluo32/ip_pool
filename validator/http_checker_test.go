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
	isValid, err := c.CheckRawProxy("www.douban.com", model.IPItem{IP: "117.69.51.122", Port: 42888})
	assert.Nil(t, err)
	assert.Equal(t, isValid, true)
}

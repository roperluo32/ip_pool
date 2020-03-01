package httpchecker

import (
	"ip_proxy/component/config"
	"ip_proxy/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpCheckerBasic(t *testing.T) {
	config.Init("test.conf", "../../..")
	c := HTTPChecker{}
	isValid, err := c.CheckProxyValid("www.douban.com", model.IPItem{IP: "117.69.51.122", Port: 42888})
	assert.NotNil(t, err)
	assert.Equal(t, isValid, false)
}

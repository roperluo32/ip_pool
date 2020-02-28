package validator

import (
	"ip_proxy/model"
)

type MockChecker struct {
}

func (mc *MockChecker) CheckRawProxy(domain string, proxy model.IPItem) (bool, error) {
	if proxy.IP == "127.0.0.1" {
		return true, nil
	}

	if proxy.IP == "192.168.0.2" {
		return false, nil
	}
	if proxy.IP == "17.291.192.6" {
		return true, nil
	}
	if proxy.IP == "39.27.92.38" {
		return false, nil
	}
	if proxy.IP == "58.19.0.23" {
		return false, nil
	}

	return false, nil
}

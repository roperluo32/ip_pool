package redisstorage

import (
	"fmt"
	"ip_proxy/model"

	"strconv"
	"strings"
)

func ipItemToString(proxy model.IPItem) string {
	return fmt.Sprintf("%s:%d", proxy.IP, proxy.Port)
}

func stringToIPItem(key string) model.IPItem {

	parts := strings.Split(key, ":")
	port, _ := strconv.Atoi(parts[1])
	return model.IPItem{
		IP:   parts[0],
		Port: port,
	}
}

func validProxyKey(domain string) string {
	return "valid_" + domain
}

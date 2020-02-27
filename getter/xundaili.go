package getter

import (
	"encoding/json"
	"fmt"

	"ip_proxy/config"
	"ip_proxy/producer"

	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/errors"
)

var once sync.Once
var _xunDailiGetter XunDaiLi

// XunDailiResponse 请求代理的回复
type XunDailiResponse struct {
	ErrorCode string `json:"ERRORCODE"`
	Result    []struct {
		IP   string `json:"ip"`
		Port string `json:"port"`
	} `json:"RESULT,omitempty"`
}

// XunDaiLi 讯代理
type XunDaiLi struct {
	// 请求间隔
	interval time.Duration
	// 请求超时
	timeout time.Duration
	// 请求参数
	reqURL     string
	method     string
	orderNo    string
	returnType string
	count      string
	spiderID   string
}

// GetInterval 抓取间隔
func (xd *XunDaiLi) GetInterval() time.Duration {
	return xd.interval
}

// GetProxyIPs 抓取代理ip
func (xd *XunDaiLi) GetProxyIPs() ([]producer.IPItem, error) {
	req, err := http.NewRequest(xd.method, xd.reqURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "http new request fail")
	}

	query := req.URL.Query()
	query.Add("orderno", xd.orderNo)
	query.Add("spiderId", xd.spiderID)
	query.Add("returnType", xd.returnType)
	query.Add("count", xd.count)
	req.URL.RawQuery = query.Encode()

	client := http.Client{Timeout: xd.timeout}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "req xundaili fail")
	}

	var xdaliResp XunDailiResponse
	if err := json.NewDecoder(resp.Body).Decode(&xdaliResp); err != nil {
		return nil, errors.Wrap(err, "decode xundaili repsonse from http body fail")
	}
	if xdaliResp.ErrorCode != "0" {
		return nil, fmt.Errorf("xdaili response code not equal 0.response:%v", xdaliResp)
	}

	var ipItems []producer.IPItem
	for _, res := range xdaliResp.Result {
		port, _ := strconv.Atoi(res.Port)
		newItem := producer.IPItem{
			IP:   res.IP,
			Port: port,
		}
		ipItems = append(ipItems, newItem)
	}

	return ipItems, nil
}

// GetCount 获取每次抓取的数量
func (xd *XunDaiLi) GetCount() string {
	return xd.count
}

// NewXunDaiLiGetter 新建一个讯代理获取器
func NewXunDaiLiGetter() producer.ProxyGetter {
	once.Do(func() {
		_xunDailiGetter = XunDaiLi{
			method:     "GET",
			reqURL:     "http://api.xdaili.cn/xdaili-api/greatRecharge/getGreatIp",
			interval:   time.Duration(config.C.XunDaiLi.Interval) * time.Second,
			returnType: config.C.XunDaiLi.ReturnType,
			spiderID:   config.C.XunDaiLi.SpiderID,
			orderNo:    config.C.XunDaiLi.OrderNo,
			count:      config.C.XunDaiLi.Count,
			timeout:    time.Duration(config.C.XunDaiLi.Timeout) * time.Second,
		}
	})

	return &_xunDailiGetter
}

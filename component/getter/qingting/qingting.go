package qingting

import (
	"encoding/json"
	"fmt"

	"ip_proxy/component/config"
	"ip_proxy/model"

	"net/http"
	"strconv"
	"sync"
	"time"

	"ip_proxy/component/log"

	"github.com/pkg/errors"
)

var once sync.Once
var _qingTingGetter QingTing

// QingTingResponse 请求代理的回复
type QingTingResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"RESULT,omitempty"`
}

// QingTing 蜻蜓代理
type QingTing struct {
	// 请求间隔
	interval time.Duration
	// 请求超时
	timeout time.Duration
	// 请求参数
	reqURL string
}

// GetInterval 抓取间隔
func (qt *QingTing) GetInterval() time.Duration {
	return qt.interval
}

// GetProxyIPs 抓取代理ip
func (qt *QingTing) GetProxyIPs() ([]model.IPItem, error) {
	req, err := http.NewRequest("GET", qt.reqURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "http new request fail")
	}

	client := http.Client{Timeout: qt.timeout}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "req QingTing fail")
	}

	var qtdailiResp QingTingResponse
	if err := json.NewDecoder(resp.Body).Decode(&qtdailiResp); err != nil {
		return nil, errors.Wrap(err, "decode QingTing repsonse from http body fail")
	}
	if qtdailiResp.Code != 0 {
		return nil, fmt.Errorf("qing ting response code not equal 0.response:%v", qtdailiResp)
	}

	var ipItems []model.IPItem
	for _, res := range qtdailiResp.Data {
		port, _ := strconv.Atoi(res.Port)
		newItem := model.IPItem{
			IP:   res.Host,
			Port: port,
		}
		ipItems = append(ipItems, newItem)
	}

	log.Infof("get raw ips:%v", ipItems)
	return ipItems, nil
}

// NewQingTingGetter 新建一个讯代理获取器
func NewQingTingGetter() *QingTing {
	once.Do(func() {
		_qingTingGetter = QingTing{
			reqURL:   config.C.QingTing.ReqURL,
			timeout:  time.Duration(config.C.QingTing.Timeout) * time.Second,
			interval: time.Duration(config.C.QingTing.Interval) * time.Second,
		}
	})

	return &_qingTingGetter
}

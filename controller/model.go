package controller

type GetProxyReq struct {
	Count int
}

type GetProxyRsp struct {
	Proxies []string
}

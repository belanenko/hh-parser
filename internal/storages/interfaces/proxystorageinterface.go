package interfaces

import "github.com/hh-parser/internal/models/proxy"

type ProxyStorage interface {
	Add(proxies ...proxy.Proxy)
	GetProxy() proxy.Proxy
	GetFormatedProxy() string
}

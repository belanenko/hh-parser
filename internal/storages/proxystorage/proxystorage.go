package proxystorage

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/hh-parser/internal/models/proxy"
	"github.com/hh-parser/internal/utli/filereader"
	"github.com/hh-parser/internal/utli/formater"
)

func init() {
	Storage = &storage{}
}

type storage struct {
	m       sync.Mutex
	Proxies []proxy.Proxy
	index   int
}

var Storage *storage

func GetStorage() *storage {
	return Storage
}

func (s *storage) Add(proxies ...proxy.Proxy) {
	Storage.m.Lock()
	defer Storage.m.Unlock()
	for _, proxy := range proxies {
		if net.ParseIP(proxy.Host) == nil || proxy.Port <= 0 || proxy.Port > 65535 {
			log.Printf("proxy is incorrect: %s:%d", proxy.Host, proxy.Port)
			continue
		}
	}

	Storage.Proxies = append(Storage.Proxies, proxies...)
}

func (s *storage) AddFromFile(path string, schema proxy.Schema) {
	strProxies := filereader.ReadAllLines(path)

	for _, p := range strProxies {
		if p == "" {
			continue
		}
		s.Add(formater.StringToProxy(proxy.SOCKS5, p)...)
	}
}

func (s *storage) GetProxy() proxy.Proxy {
	s.m.Lock()
	defer s.m.Unlock()
	if s.index == len(s.Proxies) {
		s.index = 0
	}
	this_index := s.index
	s.index++
	if len(s.Proxies) == 0 {
		log.Fatalln("no proxy in storage")
	}
	return s.Proxies[this_index]
}

func (s *storage) GetFormatedProxy() string {
	thisProxy := s.GetProxy()
	if thisProxy.UserInfo == (proxy.UserInfo{}) {
		return fmt.Sprintf("%s://%s:%d", thisProxy.Schema, thisProxy.Host, thisProxy.Port)
	}
	return fmt.Sprintf("%s://%s:%s@%s:%d", thisProxy.Schema, thisProxy.Username, thisProxy.Password, thisProxy.Host, thisProxy.Port)
}

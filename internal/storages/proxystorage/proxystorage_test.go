package proxystorage

import (
	"testing"

	"github.com/hh-parser/internal/models/proxy"
	"github.com/stretchr/testify/assert"
)

func TestAddSingletone(t *testing.T) {
	firstStorage := Storage
	secondStorage2 := Storage

	firstStorage.Proxies = append(firstStorage.Proxies, proxy.Proxy{
		Schema: proxy.HTTP,
		Host:   "209.127.191.180",
		Port:   9279,
		UserInfo: proxy.UserInfo{
			Username: "untrxaih",
			Password: "kmff7s4ojnto",
		},
	})
	firstStorage.Proxies = append(firstStorage.Proxies, proxy.Proxy{})
	firstStorage.Proxies = append(firstStorage.Proxies, proxy.Proxy{})
	firstStorage.Proxies = append(firstStorage.Proxies, proxy.Proxy{})

	assert.EqualValues(t, len(firstStorage.Proxies), len(secondStorage2.Proxies))
	Storage = &storage{}
}

func TestGetProxy(t *testing.T) {
	this_storage := Storage
	assert.Len(t, this_storage.Proxies, 0)

	this_storage.Add(
		proxy.Proxy{
			Schema: proxy.HTTP,
			Host:   "209.127.191.180",
			Port:   9279,
			UserInfo: proxy.UserInfo{
				Username: "untrxaih",
				Password: "kmff7s4ojnto",
			},
		},

		proxy.Proxy{
			Schema: proxy.HTTP,
			Host:   "45.142.28.83",
			Port:   8094,
			UserInfo: proxy.UserInfo{
				Username: "untrxaih",
				Password: "kmff7s4ojnto",
			},
		},
	)

	assert.NotEqualValues(t, this_storage.GetProxy(), this_storage.GetProxy())
	Storage = &storage{}
}
func TestGetProxyFormated(t *testing.T) {
	this_storage := Storage
	assert.Len(t, this_storage.Proxies, 0)

	this_storage.Add(
		proxy.Proxy{
			Schema: proxy.HTTP,
			Host:   "209.127.191.180",
			Port:   9279,
			UserInfo: proxy.UserInfo{
				Username: "untrxaih",
				Password: "kmff7s4ojnto",
			},
		},

		proxy.Proxy{
			Schema: proxy.HTTP,
			Host:   "45.142.28.83",
			Port:   8094,
		},
	)

	assert.EqualValues(t, this_storage.GetFormatedProxy(), "http://untrxaih:kmff7s4ojnto@209.127.191.180:9279")
	assert.EqualValues(t, this_storage.GetFormatedProxy(), "http://45.142.28.83:8094")
	Storage = &storage{}
}

func TestAddProxy(t *testing.T) {
	this_storage := Storage

	this_storage.Add(
		proxy.Proxy{
			Schema: proxy.HTTP,
			Host:   "209.127.191.180",
			Port:   9279,
			UserInfo: proxy.UserInfo{
				Username: "untrxaih",
				Password: "kmff7s4ojnto",
			},
		},

		proxy.Proxy{
			Schema: proxy.HTTP,
			Host:   "45.142.28.83",
			Port:   8094,
			UserInfo: proxy.UserInfo{
				Username: "untrxaih",
				Password: "kmff7s4ojnto",
			},
		},
	)

	assert.Len(t, this_storage.Proxies, 2)
	Storage = &storage{}
}

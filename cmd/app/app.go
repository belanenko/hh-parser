package main

import (
	"fmt"

	"github.com/hh-parser/internal/models/proxy"
	"github.com/hh-parser/internal/run/onethreed"
	"github.com/hh-parser/internal/storages/proxystorage"
	"github.com/hh-parser/internal/storages/vacancystorage"
)

func main() {
	this_proxyStorage := proxystorage.GetStorage()
	this_proxyStorage.Add(
		proxy.Proxy{
			Schema: proxy.HTTP,
			Host:   "p.webshare.io",
			Port:   80,
			UserInfo: proxy.UserInfo{
				Username: "untrxaih-rotate",
				Password: "kmff7s4ojnto",
			},
		},
	)

	this_vacancystorage := vacancystorage.Storage

	onethreed.RunOneThreed(50000000, 3, this_vacancystorage, this_proxyStorage)
	fmt.Println(this_vacancystorage.Len())

	for _, v := range this_vacancystorage.PushCount(5) {
		fmt.Println(v.Title)
	}
	fmt.Println(this_vacancystorage.Len())

}

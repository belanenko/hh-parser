package main

import (
	"fmt"
	"log"

	"github.com/hh-parser/internal/models/proxy"
	"github.com/hh-parser/internal/run/multithread"
	"github.com/hh-parser/internal/storages/proxystorage"
	"github.com/hh-parser/internal/storages/vacancystorage"
	"github.com/hh-parser/internal/utli/filereader"
	"github.com/hh-parser/internal/utli/formater"
)

func main() {
	this_proxyStorage := proxystorage.GetStorage()
	strProxies := filereader.ReadAllLines("/home/tim/code/github.com/belanenko/hh-parser/assets/socks5.txt")

	for _, p := range strProxies {
		if p == "" {
			continue
		}
		this_proxyStorage.Add(formater.StringToProxy(proxy.SOCKS5, p)...)
	}

	log.Printf("Was loaded %d proxies in storage", len(this_proxyStorage.Proxies))
	this_vacancystorage := vacancystorage.Storage

	multithread.RunMultiThread(50000000, 250, this_vacancystorage, this_proxyStorage, 5)
	fmt.Println(this_vacancystorage.Len())

	for _, v := range this_vacancystorage.PushCount(500) {
		fmt.Println(v.Title)
	}
	fmt.Println(this_vacancystorage.Len())

}

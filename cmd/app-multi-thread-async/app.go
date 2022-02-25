package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/hh-parser/internal/models/proxy"
	"github.com/hh-parser/internal/run/multithreadasync"
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

	var wg sync.WaitGroup
	var done bool
	wg.Add(1)
	go multithreadasync.Run(50000000, 5000, this_vacancystorage, this_proxyStorage, 5, wg, &done)
	fmt.Println(this_vacancystorage.Len())

	for {
		if done && this_vacancystorage.Len() == 0 {
			break
		}

		pushed := this_vacancystorage.PushCount(500)
		for _, p := range pushed {
			fmt.Printf("id: %d\t| %s\n", p.Identifier.Value, p.Title)
		}

		time.Sleep(1 * time.Second)
	}

	fmt.Println(this_vacancystorage.Len())

}
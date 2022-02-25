package main

import (
	"flag"
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
	numGorutines := flag.Int("t", 1, "count of threads")
	proxyFilePath := flag.String("pfp", "", "proxy file path")
	indexStart := flag.Int("is", 0, "Start parsing index")
	indexCount := flag.Int("ic", 1, "Count vacancy to parsing")
	flag.Parse()

	this_proxyStorage := proxystorage.GetStorage()
	strProxies := filereader.ReadAllLines(*proxyFilePath)

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
	go multithreadasync.Run(*indexStart, *indexCount, this_vacancystorage, this_proxyStorage, *numGorutines, wg, &done)
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

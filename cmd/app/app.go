package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/hh-parser/internal/models/proxy"
	worker "github.com/hh-parser/internal/run/worker"
	"github.com/hh-parser/internal/storages/proxystorage"
	"github.com/hh-parser/internal/storages/vacancystorage"
	"github.com/hh-parser/internal/utli/filereader"
	"github.com/hh-parser/internal/utli/formater"
	"github.com/hh-parser/pkg/flags"
)

func main() {
	var numGorutines, proxyFilePath, indexStart, indexCount = flags.Set()

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
	go worker.Run(*indexStart, *indexCount, this_vacancystorage, this_proxyStorage, *numGorutines, wg, &done)
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

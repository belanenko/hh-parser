package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/hh-parser/internal/models/proxy"
	worker "github.com/hh-parser/internal/run/worker"
	"github.com/hh-parser/internal/storages/proxystorage"
	"github.com/hh-parser/internal/storages/vacancystorage"
	"github.com/hh-parser/pkg/dbclick"
	"github.com/hh-parser/pkg/dbclick/clickconfig"
	"github.com/hh-parser/pkg/flags"
)

func main() {
	flagsConf := flags.Set()

	clickConfig := clickconfig.ReadConfig(flagsConf.ClickConfigPath)
	dbclick.Ping(clickConfig)

	this_proxyStorage := proxystorage.GetStorage()
	this_proxyStorage.AddFromFile(flagsConf.ProxyFilePath, proxy.SOCKS5)

	log.Printf("Was loaded %d proxies in storage", len(this_proxyStorage.Proxies))
	this_vacancystorage := vacancystorage.Storage

	var wg sync.WaitGroup
	var done bool
	wg.Add(1)
	go worker.Run(flagsConf.IndexStart, flagsConf.Count, this_vacancystorage, this_proxyStorage, flagsConf.NumGorutines, wg, &done)

	for {
		if done && this_vacancystorage.Len() == 0 {
			break
		}

		for _, p := range this_vacancystorage.PushCount(500) {
			fmt.Printf("id: %d\t| %s\n", p.Identifier.Value, p.Title)
		}

		time.Sleep(1 * time.Second)
	}
}

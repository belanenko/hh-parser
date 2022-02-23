package onethreed

import (
	"fmt"
	"log"

	"github.com/hh-parser/internal/models/proxy"
	"github.com/hh-parser/internal/models/vacancy"
	"github.com/hh-parser/internal/requests/hh"
	"github.com/hh-parser/internal/storages/proxystorage"
	"github.com/hh-parser/internal/storages/vacancystorage"
)

func Run(startIndex, countPages int) {
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
	baseUrl := `https://hh.ru/vacancy/`

	stopIndex := startIndex + countPages

	for thisIndex := startIndex; thisIndex < stopIndex; thisIndex++ {
		var newVacancy vacancy.Vacancy
		this_url := fmt.Sprintf("%s%d", baseUrl, thisIndex)
		statuscode, err := hh.GetVacancy(this_url, this_proxyStorage.GetFormatedProxy(), &newVacancy)
		if err != nil {
			log.Println(err)
			continue
		}
		if statuscode != 200 {
			log.Printf("status code is not 200, code: %d", statuscode)
			continue
		}

		this_vacancystorage.Pop(newVacancy)
	}

}

package multithreadasync

import (
	"fmt"
	"log"
	"sync"

	"github.com/hh-parser/internal/models/vacancy"
	"github.com/hh-parser/internal/requests/hh"
	"github.com/hh-parser/internal/storages/interfaces"
)

func Run(startIndex, countPages int, vs interfaces.VacancyStorage, ps interfaces.ProxyStorage, threadCount int, wgg sync.WaitGroup, done *bool) {
	defer wgg.Done()

	stopIndex := startIndex + countPages

	var wg sync.WaitGroup
	urls := make(chan string, threadCount)
	for i := 0; i < threadCount; i++ {
		wg.Add(1)
		go worker(&wg, urls, vs, ps)
	}

	for thisIndex := startIndex; thisIndex < stopIndex; thisIndex++ {
		this_url := fmt.Sprintf("https://hh.ru/vacancy/%d", thisIndex)
		urls <- this_url
	}

	close(urls)
	wg.Wait()
	*done = true
}

func worker(wg *sync.WaitGroup, urls chan string, vs interfaces.VacancyStorage, ps interfaces.ProxyStorage) {
	wg.Done()

	for url := range urls {
		var newVacancy vacancy.Vacancy
		statusCode, err := hh.GetVacancy(url, ps.GetFormatedProxy(), &newVacancy)
		if err != nil {
			log.Println(err)
			continue
		}
		if statusCode != 200 {
			log.Printf("status code is not 200, code: %d", statusCode)
			continue
		}

		vs.Pop(newVacancy)
	}
}

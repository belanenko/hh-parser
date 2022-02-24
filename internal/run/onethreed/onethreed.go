package onethreed

import (
	"fmt"
	"log"

	"github.com/hh-parser/internal/models/vacancy"
	"github.com/hh-parser/internal/requests/hh"
	"github.com/hh-parser/internal/storages/interfaces"
)

func RunOneThreed(startIndex, countPages int, vs interfaces.VacancyStorage, ps interfaces.ProxyStorage) {

	baseUrl := `https://hh.ru/vacancy/`

	stopIndex := startIndex + countPages

	for thisIndex := startIndex; thisIndex < stopIndex; thisIndex++ {
		var newVacancy vacancy.Vacancy
		this_url := fmt.Sprintf("%s%d", baseUrl, thisIndex)
		statuscode, err := hh.GetVacancy(this_url, ps.GetFormatedProxy(), &newVacancy)
		if err != nil {
			log.Println(err)
			continue
		}
		if statuscode != 200 {
			log.Printf("status code is not 200, code: %d", statuscode)
			continue
		}

		vs.Pop(newVacancy)
	}

}

package onethreed

import (
	"fmt"

	"github.com/hh-parser/internal/models/vacancy"
	"github.com/hh-parser/internal/requests/hh"
<<<<<<< HEAD
	"github.com/hh-parser/internal/storages/interfaces"
)

func RunOneThreed(startIndex, countPages int, vs interfaces.VacancyStorage, ps interfaces.ProxyStorage) {

=======
	"github.com/hh-parser/internal/storages/proxystorage"
)

func Run() {
	this_proxyStorage := proxystorage.GetStorage()
	this_proxyStorage.Add(
		proxy.Proxy{
			Schema: proxy.HTTP,
			Host:   "209.127.191.180",
			Port:   9279,
			UserInfo: proxy.UserInfo{
				Username: "untrxaih",
				Password: "kmff7s4ojnto",
			},
		},
	)

	// this_vacancystorage := vacancystorage.Storage
>>>>>>> a7f802053cf9c145d5eac407465ea2dbda66a14b
	baseUrl := `https://hh.ru/vacancy/`
	startIndex := 50000050
	stopIndex := startIndex + 15

	for thisIndex := startIndex; thisIndex < stopIndex; thisIndex++ {
		var newVacancy vacancy.Vacancy
		this_url := fmt.Sprintf("%s%d", baseUrl, thisIndex)
<<<<<<< HEAD
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
=======
		hh.GetVacancy(this_url, this_proxyStorage.GetFormatedProxy(), &newVacancy)
		switch {
		case newVacancy.BaseSalary.Value.MaxValue != 0 && newVacancy.BaseSalary.Value.MinValue != 0:
			fmt.Printf("Название вакансии: %s, Зарплата от: %d до %d \n", newVacancy.Title, newVacancy.BaseSalary.Value.MinValue, newVacancy.BaseSalary.Value.MaxValue)
		case newVacancy.BaseSalary.Value.MaxValue == 0 && newVacancy.BaseSalary.Value.MinValue != 0:
			fmt.Printf("Название вакансии: %s, Зарплата от: %d  \n", newVacancy.Title, newVacancy.BaseSalary.Value.MinValue)
		case newVacancy.BaseSalary.Value.MaxValue != 0 && newVacancy.BaseSalary.Value.MinValue == 0:
			fmt.Printf("Название вакансии: %s, Зарплата до: %d  \n", newVacancy.Title, newVacancy.BaseSalary.Value.MaxValue)
		}
>>>>>>> a7f802053cf9c145d5eac407465ea2dbda66a14b
	}

}

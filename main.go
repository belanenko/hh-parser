package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/hh-parser/storage"
	"github.com/hh-parser/structs"
)

func main() {
	res, err := http.Get(`https://hh.ru/vacancy/51044602`)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	this_storage := storage.CreateInfoStorage()

	doc.Find(".row-content").Each(func(i int, s *goquery.Selection) {
		this_name := s.Find(".vacancy-title").Find(".bloko-header-1").Text()                                                  // Название вакансии
		this_payment := s.Find(".vacancy-title").Find(".bloko-header-2_lite").Text()                                          // Размер оплаты
		this_description := s.Find(".vacancy-branded-user-content").Text()                                                    // Описание вакансии
		this_company_name := s.Find(".vacancy-company-name").Text()                                                           // Компания работодатель
		this_company_stars := s.Find(".HH-ProxyExternalServiceLoader-employerReviewsSmall").Find(".bloko-text").Text()        // Количество звезд у компании
		this_company_reviewsCount := s.Find(".HH-ProxyExternalServiceLoader-employerReviewsSmall").Find(".bloko-link").Text() // Количество отзывов у компании

		if this_name != "" {
			this_storage.AddJobs(structs.Jobs{
				Name:        this_name,
				Payment:     this_payment,
				Description: this_description,

				Company: structs.Company{
					Name:         this_company_name,
					Stars:        this_company_stars,
					ReviewsCount: this_company_reviewsCount,
				},
			})
		}
	})

	fmt.Println("Count jobs in storage:", this_storage.GetCount())

	cmpns := this_storage.FindAllByCompany("перекресток")

	fmt.Printf("%+#v", cmpns[0])

	// this_jobs := &structs.Jobs{}
	// this_jobs_ptr := &this_jobs
	// doc.Find(".row-content").Each(func(i int, s *goquery.Selection) {
	// 	(*this_jobs_ptr).Name = s.Find(".vacancy-title").Find(".bloko-header-1").Text() // Название вакансии в поле структуры
	// 	fmt.Println(s.Find(".vacancy-title").Find(".bloko-header-1").Text())               // Название вакансии в терминал
	// 	this_storage.AddJobs()
	// })

}

package hh

import (
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/hh-parser/internal/models/vacancy"
	"github.com/hh-parser/internal/utli/convertor"
	"github.com/hh-parser/internal/utli/jsonparser"

	browser "github.com/EDDYCJY/fake-useragent"
)

func GetVacancy(strUrl string, strProxy string, emptyVacancy *vacancy.Vacancy) (int, error) {
	proxyURL, err := url.Parse(strProxy)
	if err != nil {
		return 0, err
	}
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}

	req, err := http.NewRequest("GET", strUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", browser.Computer())

	resp, err := client.Do(req)
	if err != nil {
		return 1, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, nil
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return 3, err
	}
	parsed := jsonparser.GetJsonFromHTML(string(bodyBytes))
	return 200, convertor.JsonToModel(parsed, emptyVacancy)
}

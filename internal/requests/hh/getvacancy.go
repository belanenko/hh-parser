package hh

import (
	"io"
	"net/http"
	"net/url"

	"github.com/hh-parser/internal/models/vacancy"
	"github.com/hh-parser/internal/utli/convertor"
	"github.com/hh-parser/internal/utli/jsonparser"
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

	resp, err := client.Get(strUrl)
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

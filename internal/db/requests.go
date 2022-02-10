package db

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Request string

const (
	GET  Request = "GET"
	POST Request = "POST"
)

func SendRequest(reqType Request, client *http.Client, config Config, queryParam string) string {
	var req *http.Request
	switch reqType {
	case GET:
		req, _ = http.NewRequest("GET", config.Host, nil)
		query := req.URL.Query()
		query.Add("database", config.DbName)
		query.Add("query", queryParam)
		req.URL.RawQuery = query.Encode()

	case POST:
		req, _ = http.NewRequest("POST", config.Host, bytes.NewBufferString(queryParam))
	}

	req.Header.Add("X-ClickHouse-User", config.User.Name)
	req.Header.Add("X-ClickHouse-Key", config.User.Password)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	return string(data)
}

func SendPost(client *http.Client, config Config, queryParam string) string {
	req, _ := http.NewRequest("POST", config.Host, bytes.NewBufferString(queryParam))
	req.Header.Add("X-ClickHouse-User", config.User.Name)
	req.Header.Add("X-ClickHouse-Key", config.User.Password)
	resp, err := client.Do(req)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("Error : %s", err)
	}

	out, _ := io.ReadAll(req.Body)
	return string(out)
}

func SendGet(client *http.Client, config Config, queryParam string) string {
	req, _ := http.NewRequest("GET", config.Host, nil)
	query := req.URL.Query()
	query.Add("database", config.DbName)
	query.Add("query", queryParam)

	req.URL.RawQuery = query.Encode()

	req.Header.Add("X-ClickHouse-User", config.User.Name)
	req.Header.Add("X-ClickHouse-Key", config.User.Password)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	return string(data)
}

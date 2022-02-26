package dbclick

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hh-parser/pkg/dbclick/clickconfig"
)

func Ping(config *clickconfig.Config) string {
	caCert, err := ioutil.ReadFile(config.DbCertPath)
	if err != nil {
		panic(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}

	req, _ := http.NewRequest("GET", config.DbHost, nil)
	query := req.URL.Query()
	query.Add("database", config.DbName)
	query.Add("query", "SHOW TABLES FROM hh1;")

	req.URL.RawQuery = query.Encode()

	req.Header.Add("X-ClickHouse-User", config.DbUsername)
	req.Header.Add("X-ClickHouse-Key", config.DbPassword)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	data, _ := ioutil.ReadAll(resp.Body)
	answer := string(data)

	return answer
}

func Send(config *clickconfig.Config, requests ...string) {
	caCert, err := ioutil.ReadFile(config.DbCertPath)
	if err != nil {
		panic(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}

	req, _ := http.NewRequest("GET", config.DbHost, nil)
	query := req.URL.Query()
	query.Add("database", config.DbName)

	for _, request := range requests {
		query.Add("query", request)
	}

	req.URL.RawQuery = query.Encode()

	req.Header.Add("X-ClickHouse-User", config.DbUsername)
	req.Header.Add("X-ClickHouse-Key", config.DbPassword)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	data, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))
}

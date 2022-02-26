package dbclick

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"

	"github.com/hh-parser/pkg/dbclick/clickconfig"
)

func Ping(config *clickconfig.Config) string {
	return Send(config, "SHOW TABLES FROM hh1;")
}

func Send(config *clickconfig.Config, request string) string {
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

	query.Add("query", request)

	req.URL.RawQuery = query.Encode()

	req.Header.Add("X-ClickHouse-User", config.DbUsername)
	req.Header.Add("X-ClickHouse-Key", config.DbPassword)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	data, _ := ioutil.ReadAll(resp.Body)

	return string(data)
}

package dbclick

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hh-parser/internal/models/vacancy"
	"github.com/hh-parser/pkg/dbclick/clickconfig"
)

func Ping(config *clickconfig.Config) {
	answ := Send(config, "SHOW TABLES FROM hh1;")
	if answ != "row_data\n" {
		log.Fatalln("db not table row_data")
	}
	log.Println("db pinged")
}

func Batch(config *clickconfig.Config, vacancies []vacancy.Vacancy) {
	// for _, v := range vacancies {

	// }
	panic("not implimented")
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

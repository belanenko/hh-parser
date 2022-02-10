package db

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
)

func CreateClient(pathCerts ...string) *http.Client {
	caCertPool := x509.NewCertPool()

	for _, cert := range pathCerts {
		caCert, err := ioutil.ReadFile(cert)
		if err != nil {
			panic(err)
		}
		caCertPool.AppendCertsFromPEM(caCert)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}
	return client
}

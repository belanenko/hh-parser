package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	for i := 0; i < 10; i++ {
		get()
	}
}

func get() {
	url, _ := url.Parse("http://p.webshare.io:9999")
	transport := &http.Transport{
		Proxy: http.ProxyURL(url),
	}
	client := http.Client{
		Transport: transport,
	}
	resp, err := client.Get("http://ip-api.com/json/")
	if err != nil {
		log.Fatalln(err)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bodyBytes))
}

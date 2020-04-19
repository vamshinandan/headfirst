package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func getResponseSize(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("length of response for", url, len(body))
	fmt.Println(time.Now())
}

func main() {
	go getResponseSize("https://google.com")
	go getResponseSize("https://fb.com")
	go getResponseSize("https://twitter.com")
	time.Sleep(5 * time.Second)
}

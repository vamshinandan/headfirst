package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type spage struct {
	URL  string
	Size int
}

func getResponseSize(url string, sizes chan spage) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "isnt working")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	sizes <- spage{URL: url, Size: len(body)}
}

func main() {
	pages := make(chan spage)
	urls := []string{"https://google.com", "https://facebook.com", "https://twitter.com"}
	for _, url := range urls {
		go getResponseSize(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Println(page.URL, page.Size)
	}
}

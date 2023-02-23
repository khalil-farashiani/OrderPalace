package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func requestURL(url string, ch chan<- string) {
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Request failed: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Failed to read response body: %v", err)
		return
	}
	ch <- string(body)
}

func requestURLs(urls []string) {
	ch := make(chan string)
	for _, url := range urls {
		go requestURL(url, ch)
	}

	results := make([]string, len(urls))
	for i := 0; i < len(urls); i++ {
		select {
		case result := <-ch:
			results[i] = result
		case <-time.After(1 * time.Second):
			results[i] = fmt.Sprintf("Request timed out for URL: %s", urls[i])
		}
	}

	if allSuccess(results) {
		for _, result := range results {
			fmt.Println(result)
		}
	} else {
		fmt.Println("Some requests failed.")
	}
}

func allSuccess(results []string) bool {
	for _, result := range results {
		if result == "" || strings.HasPrefix(result, "Request failed:") || strings.HasPrefix(result, "Request timed out") {
			return false
		}
	}
	return true
}

func main() {
	urls := []string{"https://snappfood.ir/", "https://www.google.com/", "https://www.bing.com/"}
	requestURLs(urls)
}

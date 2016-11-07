package main

import(
	"time"
	"fmt"
	"log"
	"net/http"
)

func getStatus(urls []string) <-chan string {
	statusChan := make(chan string)
	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			statusChan <- res.Status
		} (url)
	}
	return statusChan
}

func main() {
	timeout := time.After(time.Second)
	urls := []string {
		"http://example.com",
		"http://example.net",
		"http://example.org",
	}
	statusChan := getStatus(urls)
	
	LOOP:
		for {
			select {
			case status := <- statusChan:
				fmt.Println(status)
			case <- timeout:
				break LOOP
			}
		}
	
}
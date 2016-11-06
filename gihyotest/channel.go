package main

import(
	"fmt"
	"log"
	"net/http"
)

func main() {
	urls := []string {
		"http://example.com",
		"http://example.net",
		"http://example.org",
	}
	
	statusChan := make(chan string)
	
	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			
			statusChan <- url + ": " + res.Status
		} (url)
	}
	
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-statusChan)
	} // 早くレスポンスが返って来た順番なので実行ごとに結果が異なる
	
}
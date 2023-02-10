package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// var results = make(map[string]string)
	// urls := []string{
	// 	"https://www.google.com",
	// 	"https://www.youtube.com",
	// 	"https://www.airbnb.com",
	// 	"https://www.reddit.com",
	// 	"https://www.twitter.com",
	// 	"https://www.facebook.com",
	// 	"https://www.instagram.com",
	// 	"https://www.linkedin.com",
	// 	"https://www.github.com",
	// }
	// for _, url := range urls {
	// 	result := "OK"
	// 	err := hitURL(url)
	// 	if err != nil {
	// 		result = "FAILED"
	// 	}
	// 	results[url] = result
	// }
	// for url, result := range results {
	// 	fmt.Println(url, result)
	// }

	// go routines is alive during the main function working
	go nameCount("nico")
	nameCount("giwon")
}

var errRequestFailed = fmt.Errorf("request failed")

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err)
		return errRequestFailed
	}
	return nil
}

func nameCount(persion string) {
	for i := 0; i < 10; i++ {
		fmt.Println(persion, "is me", i)
		time.Sleep(time.Second)
	}
}

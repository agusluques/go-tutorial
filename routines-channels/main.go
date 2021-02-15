package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://agustinluques.com.ar",
		"http://google.com",
		"http://github.com",
		"http://golang.org",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go getStatus(url, c)
	}

	for u := range c {
		go func(l string) {
			time.Sleep(2 * time.Second)
			getStatus(l, c)
		}(u)
	}
}

func getStatus(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url, "is down!")
		c <- url
		return
	}

	fmt.Println(url, "is up!")
	c <- url
}

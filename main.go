package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://www.amazon.com",
		"http://www.golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkStatus(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkStatus(link, c)
		}(l)
	}

}

func checkStatus(l string, c chan string) {
	_, err := http.Get(l)

	if err != nil {
		fmt.Println("Down: ", l)
		c <- l
		return
	}

	fmt.Println("Up: ", l)
	c <- l
}

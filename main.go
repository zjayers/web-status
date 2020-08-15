package main

import (
	"fmt"
	"net/http"
)

type link string

func main() {
	links := []link{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://www.amazon.com",
		"http://www.golang.org",
	}

	for _, link := range links {
		go link.checkStatus()
	}


}

func (l link) checkStatus() {
	_, err := http.Get(string(l))

	if err != nil {
		fmt.Println("Down: ", l)
		return
	}

	fmt.Println("Up: ", l)

}

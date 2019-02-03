package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/alexeyco/goozzle"
)

func main() {
	u, _ := url.Parse("https://jsonplaceholder.typicode.com/posts/1")

	_, err := goozzle.Get(u).Debug(func(res *goozzle.Response) {
		req := res.Request()

		fmt.Println("Request")
		fmt.Println("=======")
		fmt.Println("")

		fmt.Println("URL:", req.URL().String())
		fmt.Println("")

		fmt.Println("Headers:")
		for key, val := range req.Headers() {
			fmt.Println(fmt.Sprintf("%s: %s", key, val))
		}
		fmt.Println("")

		fmt.Println("Response")
		fmt.Println("=======")
		fmt.Println("")

		fmt.Println("Status:", res.Status())
		fmt.Println("")

		fmt.Println("Headers:")
		for key, val := range res.Headers() {
			fmt.Println(fmt.Sprintf("%s: %s", key, val))
		}
		fmt.Println("")

		fmt.Println("Body:")
		fmt.Println(res.String())
		fmt.Println("")
	}).Do()

	if err != nil {
		log.Fatal(err)
	}
}

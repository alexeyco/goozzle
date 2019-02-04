package goozzle_test

import (
	"fmt"
	"github.com/alexeyco/goozzle"
	"log"
	"net/url"
)

func ExampleRequest_Do() {
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

func ExampleResponse_JSON() {
	type Post struct {
		ID     int    `json:"id"`
		UserID int    `json:"userId"`
		Title  string `json:"title"`
		Body   string `json:"body"`
	}

	u, _ := url.Parse("https://jsonplaceholder.typicode.com/posts/1")

	res, err := goozzle.Get(u).Do()
	if err != nil {
		log.Fatal(err)
	}

	var post Post
	err = res.JSON(&post)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(post)
}

func ExampleRequest_JSON() {
	type Post struct {
		ID     int    `json:"id"`
		UserID int    `json:"userId"`
		Title  string `json:"title"`
		Body   string `json:"body"`
	}

	u, _ := url.Parse("https://jsonplaceholder.typicode.com/posts")

	post := &Post{
		ID:     999,
		UserID: 888,
		Title:  "Some title",
		Body:   "Many hands make light work",
	}

	_, err := goozzle.Post(u).Debug(func(res *goozzle.Response) {
		req := res.Request()

		fmt.Println("Request")
		fmt.Println("=======")
		fmt.Println("")

		fmt.Println("URL:", req.URL().String())
		fmt.Println("Method:", req.Method())
		fmt.Println("")

		fmt.Println("Headers:")
		for key, val := range req.Headers() {
			fmt.Println(fmt.Sprintf("%s: %s", key, val))
		}
		fmt.Println("")

		fmt.Println("Body:")
		fmt.Println(req.String())
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
	}).JSON(&post)

	if err != nil {
		log.Fatal(err)
	}
}

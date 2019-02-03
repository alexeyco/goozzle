package main

import (
	"log"
	"net/url"

	"github.com/alexeyco/goozzle"
)

type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
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

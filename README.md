# Goozzle
PHP Guzzle flavoured HTTP client for golang

[![Travis](https://img.shields.io/travis/alexeyco/goozzle.svg)](https://travis-ci.org/alexeyco/goozzle)
[![Coverage Status](https://coveralls.io/repos/github/alexeyco/goozzle/badge.svg?branch=master)](https://coveralls.io/github/alexeyco/goozzle?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexeyco/goozzle)](https://goreportcard.com/report/github.com/alexeyco/goozzle)
[![GoDoc](https://godoc.org/github.com/alexeyco/goozzle?status.svg)](https://godoc.org/github.com/alexeyco/goozzle)
[![license](https://img.shields.io/github/license/alexeyco/goozzle.svg)](https://github.com/alexeyco/goozzle)

See also tests and examples. If you have feature requests, send issues. Enjoy.

## Basic usage

```go
package main

import (
	"log"
	"net/url"

	"github.com/alexeyco/goozzle"
)

func main() {
	u, _ := url.Parse("https://jsonplaceholder.typicode.com/posts/1")

	res, err := goozzle.Get(u).Do()
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println(res.String())
}
```

## JSON requests

```go
type requestStruct struct {
	Foo string `json:"foo"`
}

structValue := requestStruct{
	Foo: "bar",
}

res, err := Post(u).JSON(&structValue)
```

## Send form

```go
form := url.Values{}
form.Add("Foo", "Bar")

res, err := goozzle.Post(u).Form(form)
```

## JSON response

```go
type responseStruct struct {
	Foo string `json:"foo"`
}

var value responseStruct

res, _ := Get(u).Do()
res.JSON(&value)
```

## License

```
MIT License

Copyright (c) 2019 Alexey Popov

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

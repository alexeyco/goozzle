package goozzle

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	// Version goozzle version
	Version = 1
	// UserAgentChrome if you want to pretend to be Сhrome
	UserAgentChrome = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36"
	// UserAgentSafari if you want to pretend to be Safari
	UserAgentSafari = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7"
	// UserAgentFirefox if you want to pretend to be Firefox
	UserAgentFirefox = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:63.0) Gecko/20100101 Firefox/63.0"
	// UserAgentInternetExplorer if you want to pretend to be IE
	UserAgentInternetExplorer = "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; AS; rv:11.0) like Gecko"
	// UserAgentEdge if you want to pretend to be Edge
	UserAgentEdge = "Mozilla/5.0 (Windows IoT 10.0; Android 6.0.1; WebView/3.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Mobile Safari/537.36 Edge/17.17083"
)

// UserAgentDefault default goozzle user agent
var UserAgentDefault = fmt.Sprintf("Goozzle/%d", Version)

// DebugHandler handler function to debug requests
type DebugHandler func(*Response)

// New returns request with custom method
func New(method string, url *url.URL) *Request {
	header := http.Header{}
	header.Add("User-Agent", UserAgentDefault)

	return &Request{
		method: method,
		u:      url,
		header: header,
	}
}

// Get creates GET request
func Get(url *url.URL) *Request {
	return New(http.MethodGet, url)
}

// Put creates PUT request
func Put(url *url.URL) *Request {
	return New(http.MethodPut, url)
}

// Post creates POST request
func Post(url *url.URL) *Request {
	return New(http.MethodPost, url)
}

// Delete creates DELETE request
func Delete(url *url.URL) *Request {
	return New(http.MethodDelete, url)
}

func getHeaders(headers http.Header) map[string]string {
	h := map[string]string{}

	for k := range headers {
		h[k] = headers.Get(k)
	}

	return h
}

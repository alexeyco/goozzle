package goozzle

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	Version                   = 1
	UserAgentChrome           = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36"
	UserAgentSafari           = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7"
	UserAgentFirefox          = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:63.0) Gecko/20100101 Firefox/63.0"
	UserAgentInternetExplorer = "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; AS; rv:11.0) like Gecko"
	UserAgentEdge             = "Mozilla/5.0 (Windows IoT 10.0; Android 6.0.1; WebView/3.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Mobile Safari/537.36 Edge/17.17083"
)

var UserAgentDefault = fmt.Sprintf("Goozzle/%d", Version)

type DebugRequestHandler func(*Response)

func New(method string, url *url.URL) *Request {
	header := http.Header{}
	header.Add("User-Agent", UserAgentDefault)

	return &Request{
		method: method,
		u:      url,
		header: header,
	}
}

func Get(url *url.URL) *Request {
	return New(http.MethodGet, url)
}

func Put(url *url.URL) *Request {
	return New(http.MethodPut, url)
}

func Post(url *url.URL) *Request {
	return New(http.MethodPost, url)
}

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

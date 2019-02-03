package goozzle

import (
	"bytes"
	"encoding/json"
	"golang.org/x/net/publicsuffix"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type Request struct {
	method  string
	u       *url.URL
	header  http.Header
	cookies []*http.Cookie
	body    []byte
	debug   DebugRequestHandler
}

func (r *Request) String() string {
	return string(r.body)
}

func (r *Request) Method() string {
	return r.method
}

func (r *Request) URL() *url.URL {
	return r.u
}

func (r *Request) Header(key, value string) *Request {
	r.header.Add(key, value)
	return r
}

func (r *Request) Headers() map[string]string {
	return getHeaders(r.header)
}

func (r *Request) UserAgent(userAgent string) *Request {
	r.Header("User-Agent", userAgent)
	return r
}

func (r *Request) Referer(referer string) *Request {
	r.Header("Referer", referer)
	return r
}

func (r *Request) Cookie(cookie *http.Cookie) *Request {
	r.cookies = append(r.cookies, cookie)
	return r
}

func (r *Request) Debug(h DebugRequestHandler) *Request {
	r.debug = h
	return r
}

func (r *Request) Body(body []byte) (*Response, error) {
	r.body = body
	return r.Do()
}

func (r *Request) JSON(v interface{}) (*Response, error) {
	r.header.Add("Content-Type", "application/json")

	body, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return r.Body(body)
}

func (r *Request) Form(v url.Values) (*Response, error) {
	r.header.Add("Content-Type", "application/x-www-form-urlencoded")
	return r.Body([]byte(v.Encode()))
}

func (r *Request) Do() (*Response, error) {
	client, err := r.client()
	if err != nil {
		return nil, err
	}

	var buf io.Reader = nil
	if len(r.body) > 0 {
		buf = bytes.NewBuffer(r.body)
	}

	request, err := http.NewRequest(r.method, r.u.String(), buf)
	if err != nil {
		return nil, err
	}

	for k := range r.header {
		request.Header.Add(k, r.header.Get(k))
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return r.response(response)
}

func (r *Request) client() (*http.Client, error) {
	jar, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})

	if err != nil {
		return nil, err
	}

	jar.SetCookies(r.u, r.cookies)
	client := &http.Client{
		Jar: jar,
	}

	if len(r.cookies) > 0 {
		client.Jar.SetCookies(r.u, r.cookies)
	}

	return client, nil
}

func (r *Request) response(response *http.Response) (*Response, error) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	res := &Response{
		request: r,
		status:  response.StatusCode,
		headers: response.Header,
		cookies: response.Cookies(),
		body:    body,
	}

	if r.debug != nil {
		r.debug(res)
	}

	return res, nil
}

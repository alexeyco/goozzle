package goozzle

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	request *Request
	status  int
	headers http.Header
	cookies []*http.Cookie
	body    []byte
}

func (r *Response) Request() *Request {
	return r.request
}

func (r *Response) Status() int {
	return r.status
}

func (r *Response) Header(key string) string {
	return r.headers.Get(key)
}

func (r *Response) Headers() map[string]string {
	return getHeaders(r.headers)
}

func (r *Response) Cookie(key string) string {
	for _, cookie := range r.cookies {
		if cookie.Name == key {
			return cookie.Value
		}
	}

	return ""
}

func (r *Response) Cookies() []*http.Cookie {
	return r.cookies
}

func (r *Response) Body() []byte {
	return r.body
}

func (r *Response) String() string {
	return string(r.Body())
}

func (r *Response) JSON(v interface{}) error {
	return json.Unmarshal(r.body, v)
}

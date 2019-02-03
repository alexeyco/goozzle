package goozzle

import (
	"encoding/json"
	"net/http"
)

// Response HTTP-response struct
type Response struct {
	request *Request
	status  int
	headers http.Header
	cookies []*http.Cookie
	body    []byte
}

// Request returns initial request
func (r *Response) Request() *Request {
	return r.request
}

// Status returns response status code
func (r *Response) Status() int {
	return r.status
}

// Header returns response header by name
func (r *Response) Header(key string) string {
	return r.headers.Get(key)
}

// Headers returns response headers map
func (r *Response) Headers() map[string]string {
	return getHeaders(r.headers)
}

// Cookie returns response cookie by name
func (r *Response) Cookie(key string) string {
	for _, cookie := range r.cookies {
		if cookie.Name == key {
			return cookie.Value
		}
	}

	return ""
}

// Cookies returns response cookies slice
func (r *Response) Cookies() []*http.Cookie {
	return r.cookies
}

// Body returns response body
func (r *Response) Body() []byte {
	return r.body
}

// String returns response body as string
func (r *Response) String() string {
	return string(r.Body())
}

// JSON unmarshal JSON response to struct
func (r *Response) JSON(v interface{}) error {
	return json.Unmarshal(r.body, v)
}

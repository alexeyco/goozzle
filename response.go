package goozzle

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	headers http.Header
	body    []byte
}

func (r *Response) Header(key string) string {
	return r.headers.Get(key)
}

func (r *Response) Headers() map[string]string {
	return getHeaders(r.headers)
}

func (r *Response) Body() []byte {
	return r.body
}

func (r *Response) String() string {
	return string(r.body)
}

func (r *Response) JSON(v interface{}) error {
	return json.Unmarshal(r.body, v)
}

package goozzle

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"net/url"
)

func TestRequest_Cookies(t *testing.T) {
	requestCookie := &http.Cookie{
		Name: "RequestCookie",
		Value: "Some value",
	}

	responseCookie := &http.Cookie{
		Name: "ResponseCookie",
		Value: "Another value",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookies := r.Cookies()
		if len(cookies) != 1 {
			t.Errorf("Request should contain 1 cookie, %d given", len(cookies))
		}

		if cookies[0].Name != requestCookie.Name {
			t.Errorf(`Request cookie name should be "%s", "%s" given`, requestCookie.Name, cookies[0].Name)
		}

		if cookies[0].Value != requestCookie.Value {
			t.Errorf(`Request cookie value should be "%s", "%s" given`, requestCookie.Value, cookies[0].Value)
		}

		http.SetCookie(w, responseCookie)
	}))
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	if err != nil {
		t.Error(err)
	}

	res, err := Get(u).Cookie(requestCookie).Do()
	if err != nil {
		t.Error(err)
	}

	cookies := res.Cookies()
	if len(cookies) != 1 {
		t.Errorf("Response should contain 1 cookie, %d given", len(cookies))
	}

	if cookies[0].Name != responseCookie.Name {
		t.Errorf(`Response cookie name should be "%s", "%s" given`, responseCookie.Name, cookies[0].Name)
	}

	if cookies[0].Value != responseCookie.Value {
		t.Errorf(`Response cookie value should be "%s", "%s" given`, responseCookie.Value, cookies[0].Value)
	}
}

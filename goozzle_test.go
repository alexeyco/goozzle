package goozzle

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"net/url"
	)

func Test_Headers(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf(`Method should be %s, %s given`, http.MethodDelete, r.Method)
		}

		customHeader := r.Header.Get("X-Request-Header")
		userAgent := r.Header.Get("User-Agent")
		referer := r.Header.Get("Referer")

		if customHeader != "Value" {
			t.Errorf(`Custom header "X-Request-Header" should be "%s", "%s" given`, "Value", customHeader)
		}

		if userAgent != "Test" {
			t.Errorf(`User agent should be "%s", "%s" given`, "Test", userAgent)
		}

		if referer != "http://foo.bar/fizz?buz=baz" {
			t.Errorf(`Referer should be "%s", "%s" given`, "http://foo.bar/fizz?buz=baz", referer)
		}

		w.Header().Set("X-Response-Header", "Bite me")
	}))
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	if err != nil {
		t.Error(err)
	}

	res, err := Delete(u).Header("X-Request-Header", "Value").
		UserAgent("Test").
		Referer("http://foo.bar/fizz?buz=baz").
		Do()

	if err != nil {
		t.Error(err)
	}

	responseHeader := res.Header("X-Response-Header")
	if responseHeader != "Bite me" {
		t.Errorf(`Response header should be "%s", "%s" given`, "Bite me", responseHeader)
	}
}

func Test_Cookies(t *testing.T) {
	requestCookie := &http.Cookie{
		Name: "RequestCookie",
		Value: "Some value",
	}

	responseCookie := &http.Cookie{
		Name: "ResponseCookie",
		Value: "Another value",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf(`Method should be %s, %s given`, http.MethodGet, r.Method)
		}

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

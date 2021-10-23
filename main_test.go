package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHello(t *testing.T) {
	w := httptest.NewRecorder()
	r := http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme:  "http",
			Host:    "localhost",
			Path:    "/hello",
			RawPath: "/hello",
		},
		Proto: "HTTP",
		Close: false,
		Host:  "localhost",
	}

	hello(w, &r)

	rStr, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		t.Errorf("Cannot read http response body: %s", err)
	}

	expectedRStr := "hello\n"
	if string(rStr) != expectedRStr {
		t.Errorf("Response is: %s, Want: %s", string(rStr), expectedRStr)
	}
}

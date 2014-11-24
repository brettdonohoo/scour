package scour

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

var ts *httptest.Server

func TS() *httptest.Server {
	if ts == nil {
		ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "<html><head></head><body><h1>Scour Test Heading</h1><a href='/about'>About</a></body></html>")
		}))
	}
	return ts
}

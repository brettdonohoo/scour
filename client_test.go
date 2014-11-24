package scour

import "testing"

func TestRequest(t *testing.T) {
	ts := TS()
	client := NewClient()
	_, err := client.Request("GET", ts.URL)
	if err != nil {
		t.Error(err)
	}
}

func TestRequestBad(t *testing.T) {
	client := NewClient()
	_, err := client.Request("GET", "http://localhost:666")
	if err == nil {
		t.Error(err)
	}
}

func TestGetRequest(t *testing.T) {
	ts := TS()
	client := NewClient()
	client.Request("GET", ts.URL)
	if client.req.Method != "GET" {
		t.Error("GET request not set")
	}
}

func TestPostRequest(t *testing.T) {
	ts := TS()
	client := NewClient()
	client.Request("POST", ts.URL)
	if client.req.Method != "POST" {
		t.Error("POST request not set")
	}
}

func TestSetProxy(t *testing.T) {
}

func TestSetProxyBad(t *testing.T) {
}

func TestClick(t *testing.T) {
	ts := TS()
	client := NewClient()
	crawler, _ := client.Request("GET", ts.URL)
	link, _ := crawler.SelectLink("About")
	_, err := client.Click(link)
	if err != nil {
		t.Error(err)
	}
}

func TestClickBad(t *testing.T) {
	ts := TS()
	client := NewClient()
	crawler, _ := client.Request("GET", ts.URL)
	link, _ := crawler.SelectLink("Bad Link")
	_, err := client.Click(link)
	if err == nil {
		t.Error(err)
	}
}

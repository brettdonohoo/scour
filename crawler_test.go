package scour

import "testing"

func TestSelectLinkExisting(t *testing.T) {
	ts := TS()
	client := NewClient()
	crawler, _ := client.Request("GET", ts.URL)
	_, err := crawler.SelectLink("About")
	if err != nil {
		t.Error(err)
	}
}

func TestSelectLinkNonExisting(t *testing.T) {
	ts := TS()
	client := NewClient()
	crawler, _ := client.Request("GET", ts.URL)
	_, err := crawler.SelectLink("Contact")
	if err == nil {
		t.Error(err)
	}
}

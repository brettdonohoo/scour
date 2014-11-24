package scour

import (
	"errors"
	"net/http"

	"github.com/puerkitobio/goquery"
)

// The Crawler type wraps goquery's
// document type - adding new convenience functions
type Crawler struct {
	*goquery.Document
}

// Create the new crawler type from the Response
func NewCrawler(resp *http.Response) *Crawler {
	doc, _ := goquery.NewDocumentFromResponse(resp)
	return &Crawler{doc}
}

// Look for specific anchor tag within document
func (c *Crawler) SelectLink(text string) (l *Link, err error) {

	links := []*Selection{}

	c.Find("a").Each(func(i int, s *goquery.Selection) {
		if s.Text() == text {
			links = append(links, &Selection{s, c.Url})
			// only 1 link for now
			return
		}
	})

	if len(links) == 0 {
		return nil, errors.New("Couldn't find link")
	}

	return &Link{links[0].Nodes[0], "GET", links[0].current}, nil
}

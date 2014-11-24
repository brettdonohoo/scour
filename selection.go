package scour

import (
	"net/url"

	"code.google.com/p/go.net/html"
	"github.com/puerkitobio/goquery"
)

// Link is used for holding link data
type Link struct {
	node   *html.Node
	method string
	url    *url.URL
}

func (l *Link) Method() string {
	return l.method
}

// Create the new path for Link
// TODO: Flesh this out
func (l *Link) Url() string {
	l.url.Path = l.node.Attr[0].Val
	return l.url.String()
}

// Selection embeds goquery's Selection object
// to provide a current URL
type Selection struct {
	*goquery.Selection

	// make this better?
	current *url.URL
}

// provides a link extraction function
package links

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML doc.
func Extract(url string) ([]string, error) {
	// assign value of HTTP GET request to resp
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	// if any status code other than okay, return
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	// attempt to part the response's body and assign to doc variable
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	// create an empty slice to store discovered links
	var links []string
	visitNode := func(n *html.Node) {
		// if node is an element of type a
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				// pass any attributes that aren't links
				if a.Key != "href" {
					continue
				}
				// parse the link
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				// if link is good, append to the slice we created at line 29
				links = append(links, link.String())
			}
		}
	}
	// note that nil is used as the "post" function is optional
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

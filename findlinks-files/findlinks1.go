// Findlinks1 prints the links in an HTML document read from stan input.
// usage: ../urlfetch-files/fetch https://www.awebaddress.org | ./findlinks1
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

// visit appends to links each link found in n and returns the result
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" || n.Data == "img" || n.Data == "script" || n.Data == "link" {
		for _, a := range n.Attr {
			if a.Key == "href" || a.Key == "src" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprint(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

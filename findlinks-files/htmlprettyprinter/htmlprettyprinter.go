// forEachNode calls teh functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
package htmlprettyprinter

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

var depth int

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	pre(n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	post(n)
}

func pre(node *html.Node) {
	if node.Type == html.ElementNode {
		if node.Data == "img" {
			fmt.Printf("%*s<%s/>\n", depth*2, "", node.Data)
		} else {
			fmt.Printf("%*s<%s>\n", depth*2, "", node.Data)
		}
		depth++
	}
	for _, attr := range node.Attr {
		fmt.Printf("%*s%s: %s\n", depth*2-1, "", attr.Key, attr.Val)
	}
}

func post(node *html.Node) {
	if node.Type == html.ElementNode {
		depth--
		if node.Data != "img" {
			fmt.Printf("%*s</%s>\n", depth*2, "", node.Data)
		}
	}
}

func Print(htmlString string) {
	if depth != 0 {
		depth = 0
	}
	doc, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		fmt.Printf("tried parsing htmlString: %v", err)
	}
	forEachNode(doc, pre, post)
}

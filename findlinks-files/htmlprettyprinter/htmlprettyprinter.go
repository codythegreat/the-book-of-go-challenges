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

func pre(node *html.Node) {
	if node.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", node.Data)
		depth++
	}
}

func post(node *html.Node) {
	if node.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", node.Data)
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

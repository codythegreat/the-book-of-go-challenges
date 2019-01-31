package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

// added element type to elements and increments by one
func elementCounter(elems map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		if elems[n.Data] == 0 {
			elems[n.Data] = 1
		}
		elems[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		elementCounter(elems, c)
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "elementCounter: %v\n", err)
		os.Exit(1)
	}
	elements := make(map[string]int)
	elementCounter(elements, doc)
	for k, v := range elements {
		fmt.Printf("Element: %v\t\tCount: %v\n", k, v)
	}
}

// test file for the link-grabber/links package

package main

import (
	"fmt"
	"github.com/codythegreat/book/anonymous-func-files/link-grabber"
)

func main() {
	links, err := links.Extract("https://www.wikipedia.org")
	if err != nil {
		fmt.Printf("%v", err)
	}
	for i, link := range links {
		fmt.Printf("%d\t%s\n", i+1, link)
	}
}

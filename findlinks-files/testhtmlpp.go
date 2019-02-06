// simple file to test htmlprettyprint.go package
package main

import (
	"github.com/codythegreat/book/findlinks-files/htmlprettyprinter"
)

func main() {
	nodeString := `<div id='container'><h1>Hello</h1><p>World</p></div>`
	htmlprettyprinter.Print(nodeString)
}

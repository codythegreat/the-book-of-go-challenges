// tests package tempconv

package main

import (
	"fmt"

	tempconv "github.com/codythegreat/go-book-stuff/tempconv-package"
)

func main() {
	c := tempconv.Celsius(100)
	fmt.Println(c.String())
}

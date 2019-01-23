// tests package tempconv

package main

import (
	"fmt"

	tempconv "github.com/codythegreat/go-book-stuff/tempconv-package"
)

func main() {
	c := tempconv.Celsius(100)
	k := tempconv.Kelvin(0)
	fmt.Println(c.String())
	fmt.Println(k.String())
	fmt.Println(tempconv.KToC(k).String())
}

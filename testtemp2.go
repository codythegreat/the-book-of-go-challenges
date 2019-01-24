// takes in a cmd line argument and converts it to the correct temp
package main

import (
	"fmt"
	"os"
	"strconv"

	tempconv "github.com/codythegreat/go-book-stuff/tempconv-package"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		fmt.Printf("%s = %s, %s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c), k, tempconv.KToC(k))
	}
}
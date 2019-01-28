// in place slice algorithm
package main

import "fmt"

// returns a slice of only non-empty strings
// underlying array is modified during function call
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// test the function
func main() {
	data := []string{"one", "two", "three", "", "five"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
}

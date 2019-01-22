// test the output of two functions that take in different values
package main

import "fmt"

var x, y int = 1, 2

func adder(numb *int) {
	*numb++
}

func adder2(numb int) int {
	numb++
	return numb
}

func main() {
	adder(&x)
	fmt.Println(x)
	fmt.Println(adder2(x))
}

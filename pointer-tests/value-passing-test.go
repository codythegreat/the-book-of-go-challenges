// test the side effect of two functions that take in different values
package main

import "fmt"

var x int = 1

func adder(numb *int) {
	*numb++
}

func adder2(numb int) {
	numb++
}

func main() {
	fmt.Printf("Orig x value: %d\n", x)
	adder(&x)
	fmt.Printf("x value after adder: %d\n", x)
	adder2(x)
	fmt.Printf("x value after adder2: %d\n", x)
}

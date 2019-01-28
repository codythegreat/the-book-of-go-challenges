package main

import "fmt"

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(z)
	if zlen <= cap(x) {
		// since there is room to grow, extend the slice
		z = x[:zlen]
	} else {
		// no room to grow, allocate a new array
		// Grow by doubling, for amortized linear complexity
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // build in funciton; copies x src slice to z dst slice
	}
	copy(z[len(x):], y)
	return z
}

func main() {
	fmt.Printf("%v\n", appendInt([]int{1}, 40, 50, 44, 8))
}

// prints its command line args
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	// when assigning range to two values, you get the index and the value at that index
	// to get away with not using i, we give it a black identifier _
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	end := time.Now()
	result := end.Sub(start)
	fmt.Println(result)
}

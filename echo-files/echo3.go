package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// for i, v := range os.Args[1:] {
	// 	fmt.Printf("%d: %s\n", i+1, v)
	// }

	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	end := time.Now()
	result := end.Sub(start)
	fmt.Println(result)

	// as an alternative we can do this without special formatting
	// fmt.Println(os.Args[1:])
}

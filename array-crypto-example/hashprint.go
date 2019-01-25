// prints the hash value (default SHA256) of a given input
package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"os"
)

var sha384 = flag.Bool("384", false, "prints a hash using SHA384")
var sha512 = flag.Bool("512", false, "prints a hash using SHA512")

func main() {
	if *sha384 == true {
		fmt.Printf("%d\n", 1+1)
	} else if *sha512 == true {
		fmt.Printf("%d\n", 1+1)
	} else {
		fmt.Printf("%x", sha256.Sum256([]byte(os.Args[1])))
	}
}

func printHash(s string) {

}

// returns a boolean for if two given strings are anagrams of eachother
package main

import (
	"fmt"
	"strings"
)

var anaOne string
var anaTwo string

func main() {
	anaOne = "iceman"
	anaTwo = "cinema"

	fmt.Println(testAna(anaOne, anaTwo))
}

func testAna(word1, word2 string) bool {
	if len(word1) == len(word2) {
		for _, l := range word1 {
			if !strings.Contains(word2, string(l)) {
				return false
			}
		}
	}
	return true
}

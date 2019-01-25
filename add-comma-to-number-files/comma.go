// for every three digits in a number, add a comma
package main

import (
	"bytes"
	"fmt"
	"strings"
)

// func comma(s string) string {
//	n := len(s)
//	if n <= 3 {
//		return s
//	}
//	return comma(s[:n-3]) + "," + s[n-3:]
//}

var s = "1234.789"

func main() {
	var buf bytes.Buffer
	if strings.Contains(s, "-") {
		buf.WriteRune('-')
		s = s[1:]
	}
	dotIndex := strings.Index(s, ".")
	var result string
	if dotIndex == -1 {
		result = addCommas(len(s))
	} else {
		result = addCommas(dotIndex)
	}
	buf.WriteString(result)
	fmt.Println(buf.String())
}

func addCommas(lastIndex int) string {
	result := s[lastIndex:]
	for lastIndex >= 4 {
		result = "," + s[lastIndex-3:lastIndex] + result
		lastIndex -= 3
	}
	result = s[:lastIndex] + result
	return result
}

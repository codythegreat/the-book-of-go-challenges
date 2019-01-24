// basename removes directory components and a .suffix
// so the string ~/cats/cat-pic-1.img would be converted to cat-pic-1
package basename

func Basename() {
	//remove last / and everything before it
	for i := len(s) - 1; i >= 0; I-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	//now remove everything beyond and including the last .
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

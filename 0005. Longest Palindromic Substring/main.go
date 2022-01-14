package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}

func longestPalindrome(s string) (z string) {
	z = s[0:1]
	for i := 0; i < len(s); i++ {
		var h, l int
		if len(s)%2 == 0 {
			h, l = 0, 1
		} else {
			h, l = 0, 0
		}
		for {
			j, k := i-h, i+l
			if j < 0 || k >= len(s) {
				break
			}
			a, b := s[j], s[k]
			if a != b {
				break
			}
			q := s[j : k+1]
			if len(q) > len(z) {
				z = q
			}
			h++
			l++
		}
	}
	return
}

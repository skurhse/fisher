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
		indices := [2][2]int{{0, 1}, {0, 0}}
		for n := 0; n < len(indices); n++ {
			h, l := indices[n][0], indices[n][1]
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
	}
	return
}

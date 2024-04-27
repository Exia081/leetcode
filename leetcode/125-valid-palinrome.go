package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	lowerString := strings.ToLower(s)
	for i < j {
		if !AlphanumericCharacters(lowerString[i]) {
			i++
			continue
		}
		if !AlphanumericCharacters(lowerString[j]) {
			j--
			continue
		}

		if lowerString[i] == lowerString[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func AlphanumericCharacters(b byte) bool {
	if (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
		return true
	}
	return false
}

func main() {
	s := "as  d,d,s  a"
	fmt.Println(isPalindrome(s))
}

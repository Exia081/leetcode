package main

import "fmt"

func isValid(s string) bool {
	stack := make([]byte, 0)

	sByte := []byte(s)

	for i := 0; i < len(sByte); i++ {
		if InArray(sByte[i], []byte{'[', '(', '{'}) {
			stack = append(stack, sByte[i])
		} else if len(stack) > 0 && match(stack[len(stack)-1], sByte[i]) {
			stack = stack[0 : len(stack)-1]
		} else {
			return false
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func InArray(c byte, list []byte) bool {
	for i := 0; i < len(list); i++ {
		if c == list[i] {
			return true
		}
	}
	return false
}

func match(left, right byte) bool {
	if left == '(' && right == ')' {
		return true
	}
	if left == '[' && right == ']' {
		return true
	}
	if left == '{' && right == '}' {
		return true
	}
	return false
}

/*
// beautiful answer

	func isValid(s string) bool {
		if len(s) == 0 || len(s) % 2 == 1 {
			return false
		}

		m := map[rune]rune{
			'(': ')',
			'[': ']',
			'{': '}',
		}
		stack := []rune{}

		for _, r := range s {
			if _, ok := m[r]; ok {
				stack = append(stack, r)
			} else if len(stack) == 0 || m[stack[len(stack)-1]] != r {
				return false
			} else {
				stack = stack[:len(stack) - 1]
			}
		}
		return len(stack) == 0
	}
*/
func main() {
	str := "]"
	fmt.Println(isValid(str))
}

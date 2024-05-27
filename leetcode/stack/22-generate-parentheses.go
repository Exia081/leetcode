package main

import "fmt"

/*
回溯法
1, only add open parenthesis if closed < open
2, only add a closing parenthesis if closed < open
3, valid if open == closed == n
*/
func generateParenthesis(n int) []string {
	result := make([]string, 0)
	backtracking(&result, 0, 0, n, "")
	return result
}

func backtracking(result *[]string, open, close, n int, current string) {
	if open == close && close == n {
		*result = append(*result, current)
	}
	if open < n {
		current += "("
		backtracking(result, open+1, close, n, current)
		current = current[:len(current)-1]
	}
	if close < open {
		current += ")"
		backtracking(result, open, close+1, n, current)
		current = current[:len(current)-1]
	}
}

func main() {
	result := generateParenthesis(3)
	fmt.Println(result)
}

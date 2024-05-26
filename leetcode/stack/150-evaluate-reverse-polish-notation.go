package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	nums := make([]int, 0)

	for i := 0; i < len(tokens); i++ {
		number, err := strconv.Atoi(tokens[i])
		if err == nil {
			nums = append(nums, number)
		} else {
			newNum := Calculate(nums[len(nums)-2], nums[len(nums)-1], tokens[i])
			nums = nums[:len(nums)-2]
			nums = append(nums, newNum)
		}
	}
	return nums[0]
}

func Calculate(a, b int, calc string) int {
	switch calc {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return 0
}

func main() {
	token := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(token))
	token = []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(token))
	token = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(token))
}

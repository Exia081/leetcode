package main

/*
思路：用栈保存当前还没有找到比自己大的元素，一旦找到，开始循环出栈
*/

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := make([]int, 0)

	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[index] = i - index
		}
		stack = append(stack, i)
	}
	return result
}

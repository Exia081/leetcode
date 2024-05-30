package main

import "fmt"

/*
分析：
每一个矩阵的大小到底由什么要素决定？
1，由构成这个矩阵的 直方元素最低高度 和 包含最多的直方元素 决定
2，由此可知，遍历直方图的时候，从当前元素开始，往左和往右寻找比它矮的元素，即可确定由当前元素构成的最大矩形
3，这个方式的时间复杂度是 O（ n^2 ) ,假设直方图所有元素都是相同高度，则所有元素都遍历了 n 遍。

优化：
1，我们遍历的时候，要确保选中元素是往左，往右遍历，它都是最矮的元素，不应该有比它高的元素，比它高的元素肯定能够合并进由它构成的更大的矩阵。
2，换个说法就是，我们从左右两侧分别维护一个递增的单调栈，通过维护单调栈，我们可以知道当前元素可以扩张的最大边界
3，假设边界位置的值未-1，也就是边界的值肯定就是单调栈的底部，在维护单调栈的过程中，栈顶元素即当前元素最终可以扩张的最大边界

*/

func largestRectangleArea(heights []int) int {

	singleStack := make([]int, 0)

	leftBoarderArr := make([]int, len(heights))
	rightBoarderArr := make([]int, len(heights))

	for i := 0; i < len(heights); i++ {
		for len(singleStack) > 0 && heights[singleStack[len(singleStack)-1]] >= heights[i] {
			singleStack = singleStack[:len(singleStack)-1]
		}

		if len(singleStack) == 0 {
			leftBoarderArr[i] = -1
		} else {
			leftBoarderArr[i] = singleStack[len(singleStack)-1]
		}
		singleStack = append(singleStack, i)
	}

	singleStack = []int{}
	for i := len(heights) - 1; i >= 0; i-- {
		for len(singleStack) > 0 && heights[singleStack[len(singleStack)-1]] >= heights[i] {
			singleStack = singleStack[:len(singleStack)-1]
		}
		if len(singleStack) == 0 {
			rightBoarderArr[i] = len(heights)
		} else {
			rightBoarderArr[i] = singleStack[len(singleStack)-1]
		}
		singleStack = append(singleStack, i)
	}

	maxArea := 0

	for i := 0; i < len(heights); i++ {
		area := (rightBoarderArr[i] - leftBoarderArr[i] - 1) * heights[i]
		maxArea = max(maxArea, area)
	}
	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	heights := []int{6, 7, 5, 2, 4, 5, 9, 3}
	fmt.Println(largestRectangleArea(heights))

	heights = []int{2, 4}
	fmt.Println(largestRectangleArea(heights))
}

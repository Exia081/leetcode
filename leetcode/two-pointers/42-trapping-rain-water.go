package main

import "fmt"

/*
思路分析
这题也是用双指针的来解决，但与传统的双指针不通，它的前进方向并不固定，左高则右前进，右高则左前进
*/
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	//water  =  min( leftMaxHeight,rightMaxHeight) - height i

	low := 0
	high := len(height) - 1
	water := 0
	leftMaxHeight := height[low]
	rightMaxHeight := height[high]

	for low < high {
		if leftMaxHeight < rightMaxHeight {
			low++
			leftMaxHeight = max(leftMaxHeight, height[low])
			water += leftMaxHeight - height[low]
		} else {
			high--
			rightMaxHeight = max(rightMaxHeight, height[high])
			water += rightMaxHeight - height[high]
		}
	}
	return water
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}

func main() {
	//height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	//fmt.Println(trap(height))

	height := []int{4, 2, 3}
	fmt.Println(trap(height))
}

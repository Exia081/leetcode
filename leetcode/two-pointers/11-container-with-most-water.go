package main

/*
思路，采用双指针，每一次移动，宽减少，高可能增加，只需按哪边矮就移动那边即可
*/

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}
	low := 0
	high := len(height) - 1

	maxVal := min(height[low], height[high]) * (high - low)

	for low < high {
		val := min(height[low], height[high]) * (high - low)
		if val > maxVal {
			maxVal = val
		}

		if height[low] > height[high] {
			high--
		} else {
			low++
		}
	}
	return maxVal
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

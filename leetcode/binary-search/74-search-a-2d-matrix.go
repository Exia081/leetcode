package main

import (
	"fmt"
)

/*
两个思路
1，二分查找行，再二分查找列
2，直接将整个矩阵当做数组进行二分查找
*/

func binarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (high + low) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func searchMatrix(matrix [][]int, target int) bool {

	low := 0
	high := len(matrix) - 1

	for low <= high {
		mid := (low + high) / 2
		if target >= matrix[mid][0] && target <= matrix[mid][len(matrix[mid])-1] {
			index := binarySearch(matrix[mid], target)
			if index != -1 {
				return true
			} else {
				return false
			}
		}

		if target < matrix[mid][0] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return false
}

func main() {
	matrix := [][]int{
		{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60},
	}
	fmt.Println(searchMatrix(matrix, 3))
}

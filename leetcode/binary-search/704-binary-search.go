package main

import "fmt"

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

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}

	fmt.Println(binarySearch(nums, -1))
	fmt.Println(binarySearch(nums, 12))
	fmt.Println(binarySearch(nums, 5))
	fmt.Println(binarySearch(nums, 100))
}

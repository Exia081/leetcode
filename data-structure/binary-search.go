package main

import "fmt"

func BinarySearch(arr []int, target int) (index int) {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := (end + start) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}

func binarySearch(list []int, item int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]

		if guess == item {
			return mid
		}

		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1 // 没有找到
}

func main() {
	arr := []int{1}
	fmt.Println(BinarySearch(arr, 1))
	fmt.Println(BinarySearch(arr, 0))

	arr = []int{1, 2, 3, 4, 5}
	fmt.Println(BinarySearch(arr, 1))
	fmt.Println(BinarySearch(arr, 5))
	fmt.Println(BinarySearch(arr, 3))
	fmt.Println(BinarySearch(arr, 0))
}



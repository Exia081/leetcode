package main

import "fmt"

func countingSort(arr []int, max int) []int {
	counts := make([]int, max+1)     // 初始化计数数组
	results := make([]int, len(arr)) // 初始化结果数组

	// 计数
	for _, num := range arr {
		counts[num]++
	}

	// 累加, 这样可以让最后一个元素直接对齐到数组的最后一个位置上
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}

	// 排序
	for i := len(arr) - 1; i >= 0; i-- {
		num := arr[i]
		counts[num]-- // 减一表示这个元素已经被排序
		results[counts[num]] = num
	}

	return results
}

// 该算法只适用于正整数排序，如果要负整数排序，需要找到最小值，再进行偏移操作

func main() {
	arr := []int{4, 3, 2, 5, 1, 3}
	fmt.Println("Original array:", arr)
	arr = countingSort(arr, 5) // 5 是数组中的最大值
	fmt.Println("Sorted array:", arr)
}

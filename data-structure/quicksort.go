package main

import "fmt"

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	povit := arr[start]
	i := start + 1
	j := end

	for i <= j {
		if arr[i] < povit {
			i++
		} else {
			arr[i], arr[j] = arr[j], arr[i]
			j--
		}
	}

	arr[start], arr[j] = arr[j], arr[start]

	quickSort(arr, start, j)
	quickSort(arr, j+1, end)
}

func QuickSort(arr []int) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr
}

func main() {
	arr := []int{5, 2, 6, 3, 1, 4}
	fmt.Println(QuickSort(arr))
}

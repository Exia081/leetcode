package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	sorted := make([]int, 0, len(left)+len(right))
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			sorted = append(sorted, left[0])
			left = left[1:]
		} else {
			sorted = append(sorted, right[0])
			right = right[1:]
		}
	}

	if len(left) > 0 {
		sorted = append(sorted, left...)
	} else if len(right) > 0 {
		sorted = append(sorted, right...)
	}

	return sorted
}

func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	result := mergeSort(arr)
	fmt.Println("Sorted array is", result)
}

func mergeSortWithoutRecursive(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	aux := make([]int, len(arr))
	copy(aux, arr)

	for sz := 1; sz < len(arr); sz += sz {
		for low := 0; low < len(arr)-sz; low += sz + sz {
			mid := low + sz - 1
			high := min(low+sz+sz-1, len(arr)-1)
			mergeWithoutRecursive(aux, arr, low, mid, high)
		}
	}

	return aux
}

func mergeWithoutRecursive(aux, arr []int, low, mid, high int) {
	i, j := low, mid+1

	for k := low; k <= high; k++ {
		if i > mid {
			arr[k] = aux[j]
			j++
		} else if j > high {
			arr[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			arr[k] = aux[j]
			j++
		} else {
			arr[k] = aux[i]
			i++
		}
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

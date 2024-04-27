package main

import "sort"

func threeSum(numbers []int) [][]int {

	result := make([][]int, 0)
	sort.Ints(numbers)

	for i := 0; i < len(numbers); i++ {

		low := i + 1
		high := len(numbers) - 1

		for low < high {
			if numbers[i]+numbers[low]+numbers[high] > 0 {
				high--
			} else if numbers[i]+numbers[low]+numbers[high] < 0 {
				low++
			} else {
				//相等
				result = append(result, []int{numbers[i], numbers[low], numbers[high]})

				//出清所有元素
				same := numbers[low]
				low++
				for low < len(numbers) {
					if numbers[low] == same {
						low++
					} else {
						break
					}
				}
				same = numbers[high]
				high--
				for high > i {
					if numbers[high] == same {
						high--
					} else {
						break
					}
				}
			}
		}

		//出清相同元素
		same := numbers[i]
		for i+1 < len(numbers) {
			if numbers[i+1] == same {
				i++
			} else {
				break
			}
		}
	}
	return result
}

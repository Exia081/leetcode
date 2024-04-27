package main

import "fmt"

/*
https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.

Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.

The tests are generated such that there is exactly one solution. You may not use the same element twice.

Your solution must use only constant extra space.

Example 1:

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].
Example 2:

Input: numbers = [2,3,4], target = 6
Output: [1,3]
Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].
Example 3:

Input: numbers = [-1,0], target = -1
Output: [1,2]
Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].
*/

/*
思路1
1，题目要求只能使用常数级别的额外空间，因此不能使用 hash 来做遍历
2，因为数组本身有序，我们可以对元素进行二分查找，这样就可以快速找到对应的元素是否存在
time complexity O( n log n )
space complexity O(1)
*/

func twoSumWithoutOptimization(numbers []int, target int) []int {

	for i := 0; i < len(numbers); i++ {

		//当元素已经大于目标值的一半，后续元素相加也不可能得到目标值
		if numbers[i] > target/2 {
			break
		}

		addTarget := target - numbers[i]

		if i+1 >= len(numbers) {
			break
		}

		index := BinarySearch(numbers[i+1:], addTarget)
		if index != -1 {
			return []int{i + 1, index + i + 2}
		}
	}
	return []int{}
}

func BinarySearch(numbers []int, target int) int {
	if len(numbers) == 0 {
		return -1
	}
	low := 0
	high := len(numbers) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := numbers[mid]
		if guess > target {
			high = mid - 1
		} else if guess < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

/*
思路 2
直接采用双指针遍历，因为数组有序，我们采用左加右减的方式遍历整个数组，并不会漏掉可能出现的组合
*/
func twoSum(numbers []int, target int) []int {
	if len(numbers) == 0 {
		return []int{}
	}
	low := 0
	high := len(numbers) - 1

	for low < high {
		if numbers[low]+numbers[high] > target {
			high--
		} else if numbers[low]+numbers[high] < target {
			low++
		} else {
			return []int{low + 1, high + 1}
		}
	}
	return []int{}
}

func main() {
	numbers := []int{3, 24, 50, 79, 88, 150, 345}
	target := 200
	fmt.Println(twoSum(numbers, target))
}

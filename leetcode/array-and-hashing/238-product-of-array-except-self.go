package main

import "fmt"

/*
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.



Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]


Constraints:

2 <= nums.length <= 105
-30 <= nums[i] <= 30
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.


Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)
*/

/*
思路一，保存额外两个数组，代表左乘积，右乘积
*/
func productExceptSelfWithExtraSpace(nums []int) []int {

	if len(nums) <= 1 {
		return nums
	}

	leftProduct := make([]int, len(nums))
	rightProduct := make([]int, len(nums))

	product := 1
	for i := 0; i < len(nums); i++ {
		leftProduct[i] = product
		product = product * nums[i]
	}

	product = 1
	for i := len(nums) - 1; i >= 0; i-- {
		rightProduct[i] = product
		product = product * nums[i]
	}

	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = rightProduct[i] * leftProduct[i]
	}
	return result
}

/*
Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)
思路二，不需要保存两个数组，直接将乘积保存去结果数组即可
*/
func productExceptSelf(nums []int) []int {

	if len(nums) <= 1 {
		return nums
	}

	result := make([]int, len(nums))

	product := 1
	for i := 0; i < len(nums); i++ {
		result[i] = product
		product = product * nums[i]
	}

	product = 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = product * result[i]
		product = product * nums[i]
	}

	return result
}

func main() {
	input := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(input))

	input = []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(input))

	input = []int{0}
	fmt.Println(productExceptSelf(input))
}

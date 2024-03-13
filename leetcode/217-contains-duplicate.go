package main

/*
	description : https://leetcode.com/problems/contains-duplicate/description/
*/

func containsDuplicate(nums []int) bool {
	mp := make(map[int]int)

	for _, val := range nums {
		mp[val]++

		if mp[val] > 1 {
			return true
		}
	}
	return false
}

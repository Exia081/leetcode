package main

import (
	"fmt"
)

/*
description: https://leetcode.com/problems/top-k-frequent-elements/description/
*/

/*
方案 1：
1，用 hashmap 来保存各个元素的出现次数
2，利用counting sort算法原理，进行分桶插入，将出现次数相通的元素保存在array中
time complexity : O(n)
space complexity : O(n)  if n == k still O(n)
*/
func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]int)

	for _, v := range nums {
		mp[v]++
	}

	showTimesArr := make([][]int, len(nums)+1)

	for mpk, mpv := range mp {
		showTimesArr[mpv] = append(showTimesArr[mpv], mpk)
	}

	result := make([]int, 0)

	popCnt := 0

	for i := len(showTimesArr) - 1; i >= 0; i-- {

		for _, insideVal := range showTimesArr[i] {
			result = append(result, insideVal)
			popCnt++
		}

		if popCnt == k {
			break
		}
	}

	return result
}

func topKFrequentBest(nums []int, k int) []int {
	/*
		    用计数排序方法来实现统计
			time : O(n)
			space : O(n)
	*/
	countingBucket := make([][]int, len(nums)+1)

	countingMap := make(map[int]int, 0)
	result := []int{}

	for _, v := range nums {
		countingMap[v]++
	}

	for number, count := range countingMap {
		countingBucket[count] = append(countingBucket[count], number)
	}

	for i := len(countingBucket) - 1; i >= 0 && k > 0; i-- {
		if len(countingBucket[i]) != 0 {
			result = append(result, countingBucket[i]...)
			k -= len(countingBucket[i])
		}
	}
	return result
}

/*
todo
https://leetcode.com/problems/top-k-frequent-elements/editorial/
其他方案
1，hash + heap
time: O(n log k) if k <<（远小于） n ,actually O(n), if k == n , actually O( n log n)
space: O(n + k)  actually O(n) because n >= k ,if k == n , O(n+n) still O(n)

2，还有快速排序方案，但依然没有计数排序高效
*/

func main() {
	nums := []int{3, 0, 1, 0}
	fmt.Printf("%+v", topKFrequent(nums, 1))

	nums = []int{1, 1, 1, 2, 2, 3}
	fmt.Printf("%+v", topKFrequent(nums, 2))
}

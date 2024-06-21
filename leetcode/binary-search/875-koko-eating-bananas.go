package main

import (
	"fmt"
	"math"
)

/*
分析
1，已知总香蕉数，最大吞并次数，这两个数相除，可以得出最低每次需要吞并的香蕉数，然后最大值就是当前数组里面最大的元素，这样吞并次数刚好就等于元素个数
2，因为最大值和最小值都已经知道，这样就可以进行二分查找，找到最小值

时间复杂度 O(n log k) n 代表元素个数，k代表最大的元素
空间负责度 O(1)
*/

func minEatingSpeed(piles []int, h int) int {
	sum := 0
	maxItem := -1

	for i := 0; i < len(piles); i++ {
		sum += piles[i]
		if piles[i] > maxItem {
			maxItem = piles[i]
		}
	}

	minItem := int(math.Ceil(float64(sum) / float64(h)))

	low := minItem
	high := maxItem

	result := maxItem

	//二分查找里面包含了数组遍历，因此必须大于O(n)
	//对数的最大值是maxItem

	for low <= high {
		mid := (low + high) / 2

		kCnt := 0
		for i := 0; i < len(piles); i++ {
			kCnt += int(math.Ceil(float64(piles[i]) / float64(mid)))
		}
		//fmt.Println(mid)
		if h < kCnt {
			low = mid + 1
		} else {
			high = mid - 1
			if mid < result {
				result = mid
			}
		}
	}

	return result
}

func main() {

	piles := []int{1, 4, 3, 2}
	h := 9
	fmt.Println(fmt.Sprintf("piles:%+v,h:%+v,result:%+v", piles, h, minEatingSpeed(piles, h)))

	piles = []int{3, 6, 7, 11}
	h = 8
	fmt.Println(fmt.Sprintf("piles:%+v,h:%+v,result:%+v", piles, h, minEatingSpeed(piles, h)))

	piles = []int{30, 11, 23, 4, 20}
	h = 5
	fmt.Println(fmt.Sprintf("piles:%+v,h:%+v,result:%+v", piles, h, minEatingSpeed(piles, h)))

	piles = []int{30, 11, 23, 4, 20}
	h = 6
	fmt.Println(fmt.Sprintf("piles:%+v,h:%+v,result:%+v", piles, h, minEatingSpeed(piles, h)))
}

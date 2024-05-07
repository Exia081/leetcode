package main

/*
题目的本质，不断寻找最低点，以及最低点右边比它大的元素的差值。
*/
func maxProfit(prices []int) int {
	minPrice := 999999999
	profit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}
	}
	return profit
}

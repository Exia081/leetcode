package main

import "fmt"

func longestConsecutive(nums []int) int {
	// Construct a set out of the nums array.
	numsSet := make(map[int]struct{})
	for _, n := range nums {
		numsSet[n] = struct{}{}
	}

	// The answer is stored here.
	maxSequenceLen := 0

	// Iterate through the set.
	for n := range numsSet {
		// We check if n-1 is in the set. If it is, then n is not the beginning of a sequence
		// and we go to the next number immediately.
		if _, ok := numsSet[n-1]; !ok {
			// Otherwise, we increment n in a loop to see if the next consecutive value is stored in nums.
			seqLen := 1
			for {
				if _, ok = numsSet[n+seqLen]; ok {
					seqLen++
					continue
				}
				// When the sequence is over, see if we did better than before.
				maxSequenceLen = max(seqLen, maxSequenceLen)
				break
			}
		}
	}

	return maxSequenceLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//nums := []int{100, 4, 200, 1, 3, 2}
	//fmt.Println(QuickSort(nums))
	//fmt.Println(longestConsecutive(nums))

	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	//fmt.Println(QuickSort(nums))
	fmt.Println(longestConsecutive(nums))
}

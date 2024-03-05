package main

import (
	"container/heap"
	"fmt"
)

/*
	有20个数组，每个数组有500个元素，并且是有序排列好的，
	现在如何在这20×500个数中找出排名前500的数？ 能给出 go 语言的代码实现吗
*/

type IntegerHeap []int

func (iheap IntegerHeap) Len() int { return len(iheap) }

func (iheap IntegerHeap) Less(i, j int) bool { return iheap[i] < iheap[j] }

func (iheap IntegerHeap) Swap(i, j int) { iheap[i], iheap[j] = iheap[j], iheap[i] }

func (iheap *IntegerHeap) Push(heapintf interface{}) {
	*iheap = append(*iheap, heapintf.(int))
}

func (iheap *IntegerHeap) Pop() interface{} {
	var n = len(*iheap)
	var x1 = (*iheap)[n-1]
	*iheap = (*iheap)[0 : n-1]
	return x1
}

func findTopKNumbers(arrays [][]int, k int) []int {
	var ans []int
	h := &IntegerHeap{}
	heap.Init(h)
	for i := 0; i < len(arrays); i++ {
		for j := 0; j < len(arrays[i]); j++ {
			if h.Len() < k {
				heap.Push(h, arrays[i][j])
			} else if (*h)[0] < arrays[i][j] {
				heap.Pop(h)
				heap.Push(h, arrays[i][j])
			}
		}
	}
	for h.Len() > 0 {
		ans = append(ans, heap.Pop(h).(int))
	}
	return ans
}

func main() {
	arrays := make([][]int, 20)
	for i := range arrays {
		arrays[i] = make([]int, 500)
		for j := range arrays[i] {
			arrays[i][j] = i * j // 这里只是为了示例，实际数据需要根据实际情况填入
		}
	}
	result := findTopKNumbers(arrays, 500)
	for i := range result {
		fmt.Println(result[i])
	}
}

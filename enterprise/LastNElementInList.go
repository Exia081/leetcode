package main

import "fmt"

/*
找到当前链表中 倒数第 N 个元素
*/

type Node struct {
	Val  int
	Next *Node
}

func LastNElementInList(head *Node, n int) int {
	if n <= 0 {
		return -1
	}
	var gap, next *Node
	next = head
	gap = head
	cnt := 0
	for next != nil {
		if cnt >= n {
			gap = gap.Next
		}
		cnt++
		next = next.Next
	}
	if cnt < n {
		return -1
	}
	return gap.Val
}

func main() {
	var start, prev *Node

	for i := 1; i <= 20; i++ {
		node := &Node{
			Val: i,
		}
		if start == nil {
			start = node
			prev = node
		} else {
			prev.Next = node
			prev = node
		}
	}

	//for start != nil {
	//	fmt.Println(start.Val)
	//	start = start.Next
	//}

	fmt.Println(LastNElementInList(start, 19))
}

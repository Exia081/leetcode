package main

import "fmt"

type Stack struct {
	head *Node
}

type Node struct {
	Val  int
	Next *Node
}

func (s *Stack) Push(val int) {
	n := new(Node)
	n.Val = val
	n.Next = s.head
	s.head = n
}

func (s *Stack) Pop() (val int) {
	if s.head == nil {
		return
	}

	val = s.head.Val
	s.head = s.head.Next
	return
}

func main() {
	ins := Stack{}

	ins.Push(1)
	ins.Push(2)

	fmt.Println(ins.Pop())
	fmt.Println(ins.Pop())
	fmt.Println(ins.Pop())
}

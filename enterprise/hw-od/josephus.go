package main

import "fmt"

/*
问题描述
约瑟夫问题的源头完全可以命名为“ZS游戏”。本着和谐友爱和追求本质的目的，我们把问题描述如下：

现有T个人围成一桌坐下，编号从1到T，从编号为1的人开始报数。
报数也从1开始，报到M的人离席，从离席者的下一位在座成员开始，继续从1开始报数。
复现这个过程（各成员的离席次序），或者求最后一个在座的成员编号。
*/

func JosephusRing(total, call, left int) []int {
	nowLeft := total
	nowCall := 0
	index := -1
	arr := make([]int, total)

	for {

		//注意边界 等于号
		if nowLeft <= left {
			break
		}

		index++
		if index == len(arr) {
			index = 0
		}
		for arr[index] == 1 {
			index++
			if index == len(arr) {
				index = 0
			}
		}

		if nowCall == call {
			nowCall = 1
		} else {
			nowCall += 1
		}

		//fmt.Printf("index:%d,call:%d\n", index, nowCall)

		if nowCall == call {
			arr[index] = 1
			//记得扣减数量
			nowLeft--
		}
	}
	//fmt.Println(arr)
	result := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			result = append(result, i+1)
		}
	}

	return result
}

type Node struct {
	Val  int
	Next *Node
}
type CircleLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (s *CircleLinkedList) AddNode(val int) {
	node := &Node{Val: val}
	if s.Size == 0 {
		s.Head = node
		s.Tail = node
	} else {
		node.Next = s.Head
		s.Tail.Next = node
		s.Tail = node
	}
	s.Size++
}
func (s *CircleLinkedList) RemoteNode(node *Node) {

}

func main() {
	//fmt.Println(JosephusRing(2, 3, 1))
	//fmt.Println(JosephusRing(41, 3, 2))
	fmt.Println(JosephusRing(100, 4, 3))
}

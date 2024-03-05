package main

import "fmt"

type List struct {
	head *Element
	tail *Element
}

/*
	a link list maintain head and tail pointer.

	Implement functions for link list as following:
	1, Delete(Element *elem) bool
	2, InsertAfter(Element *elem,int val) bool
*/

type Element struct {
	Val  int
	Next *Element
}

func (l *List) Delete(elem *Element) (result bool) {

	/*
		there has 2 cases:
		1, delete the head element
		2, delete element in middle
	*/

	if elem == nil || l.head == nil {
		return false
	}

	current := l.head
	//1, delete the head element
	if current.Val == elem.Val {
		l.head = l.head.Next

		if l.head == nil {
			l.tail = nil
		}
		return true
	}

	//2, delete element in middle
	for current != nil {
		if current.Next != nil && current.Next.Val == elem.Val {
			current.Next = current.Next.Next
			if current.Next == nil {
				l.tail = current
			}
			return true
		}
		current = current.Next
	}
	return false
}
func (l *List) InsertAfter(elem *Element, val int) (result bool) {
	newEl := new(Element)
	newEl.Val = val
	if elem == nil {
		newEl.Next = l.head
		l.head = newEl
		if l.tail == nil {
			l.tail = newEl
		}
		return true
	}

	current := l.head

	for current != nil {
		if current.Val == elem.Val {
			newEl.Next = current.Next
			current.Next = newEl
			if newEl.Next == nil {
				l.tail = newEl
			}
			return true
		}
		current = current.Next
	}
	return false
}

func (l List) Print() {

	content := make([]int, 0)

	cur := l.head

	for cur != nil {
		content = append(content, cur.Val)
		cur = cur.Next
	}
	fmt.Printf("head:%+v,tail:%+v,content:%+v \n", l.head, l.tail, content)
}

func main() {

	iList := List{}

	//el := Element{
	//	Val: 1,
	//}
	fmt.Println(iList.InsertAfter(&Element{
		10,
		nil,
	}, 1))
	fmt.Println(iList.InsertAfter(nil, 1))
	iList.Print()

	fmt.Println(iList.InsertAfter(nil, 2))
	iList.Print()

	fmt.Println(iList.InsertAfter(&Element{1, nil}, 3))
	iList.Print()

	fmt.Println(iList.InsertAfter(&Element{10, nil}, 3))
	iList.Print()

	fmt.Println(iList.Delete(&Element{3, nil}))
	iList.Print()
}

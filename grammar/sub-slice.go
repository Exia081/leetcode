package main

import "fmt"

func main() {

	/*
			在Go语言中，你可以使用slice[low:high]的方式来获取一个子切片，
		其中low是切片的起始索引（包括），high是结束索引（不包括）。其实high索引位置的元素是不包括在新的切片中的。
	*/

	// initialize a slice with 4 len and values
	number2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(number2) // -> [1 2 3 4]
	// create sub slices
	slice1 := number2[2:]
	fmt.Println(slice1) // -> [3 4]
	slice2 := number2[:3]
	fmt.Println(slice2) // -> [1 2 3]
	slice3 := number2[1:4]
	fmt.Println(slice3) // -> [2 3 4]

	var a = make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Printf("slice:%+v len:%+v cap:%+v", a, len(a), cap(a))

}

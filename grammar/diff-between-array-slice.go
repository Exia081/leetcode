package main

import (
	"fmt"
)

func main() {

	/*
		切片和数组的区别
		1，动态长度和固定长度
		2，引用类型和值类型
	*/

	arr := make([]int, 0, 4)

	arr = append(arr, 1, 2, 3, 4)

	fmt.Printf("slice %+v  len:%+v cap:%+v \n", arr, len(arr), cap(arr))

	change(arr)

	fmt.Printf("slice %+v  len:%+v cap:%+v \n", arr, len(arr), cap(arr))

	iArr := [3]int{1, 2, 3}
	fmt.Printf("array %+v  len:%+v cap:%+v \n", iArr, len(iArr), cap(iArr))
	changeArray(iArr)
	fmt.Printf("array %+v  len:%+v cap:%+v \n", iArr, len(iArr), cap(iArr))
}

func change(arr []int) {
	for i := 0; i < 3; i++ {
		arr[i] = 6
	}

	nArr := make([]int, 1, 256)
	for i := 0; i < 256; i++ {
		nArr = append(nArr, i)
	}

	arr = append(arr, nArr...)

	fmt.Printf("inside slice %+v  len:%+v cap:%+v \n", arr, len(arr), cap(arr))
}

func changeArray(arr [3]int) {
	for i := 0; i < len(arr); i++ {
		arr[i]++
	}
}

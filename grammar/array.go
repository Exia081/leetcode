package main

import "fmt"

func main() {

	arr := []int{1, 2, 3}
	fmt.Printf("len:%d,cap:%d\n", len(arr), cap(arr))
	fmt.Println(arr[3:])
}

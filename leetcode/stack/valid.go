package main

import "fmt"

var kMap = map[byte]byte{
	'(': ')',
}

func reverse(str string) string {
	stack := make([][]byte, 0)
	stack = append(stack, make([]byte, 0))
	bList := []byte(str)
	for i := 0; i < len(bList); i++ {
		if _, ok := kMap[bList[i]]; ok {
			stack = append(stack, make([]byte, 0))
		} else if bList[i] == ')' {
			subStr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			low := 0
			high := len(subStr) - 1
			for low < high {
				subStr[low], subStr[high] = subStr[high], subStr[low]
				low++
				high--
			}

			stack[len(stack)-1] = append(stack[len(stack)-1], subStr...)
		} else {
			stack[len(stack)-1] = append(stack[len(stack)-1], bList[i])
		}
	}

	return string(stack[len(stack)-1])
}

func main() {
	str := "adb(dac)"
	fmt.Println(reverse(str))

	str = "(ed(et(oc))el)"
	fmt.Println(reverse(str))
}

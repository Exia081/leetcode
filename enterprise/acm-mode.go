package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
1-1每行数字固定，不知道行数
*/

func GetFixedNumberInput() {
	a := 0
	b := 0
	for {
		n, _ := fmt.Scanln(&a, &b)
		if n == 0 {
			break
		} else {
			fmt.Printf("a:%d,b:%d,sum:%d\n", a, b, a+b)
		}
	}
}

/*
1-2,每行数字固定，知道行数
*/

func GetFixedNumberAndLineInput() {
	line := 0
	a, b := 0, 0
	fmt.Scan(&line)

	for i := 0; i < line; i++ {
		fmt.Scan(&a, &b)
		fmt.Println(a + b)
	}
}

/*
2-1-1,不知行数，每行数字不固定,但知道数量，知道结束符
*/

func GetRandomNumberAndLineInput() {
	LineLength := 0
	for {
		sum := 0
		fmt.Scan(&LineLength)
		if LineLength == 0 {
			break
		} else {
			arr := make([]int, LineLength)
			for i := 0; i < LineLength; i++ {
				fmt.Scan(&arr[i])
			}
			for i := 0; i < LineLength; i++ {
				sum += arr[i]
			}
		}
		fmt.Println(sum)
	}
}

/*
2-1-2,不知行数，每行数字不固定,但知道数量，结束符
*/

func GetRandomNumberAndLineInputWithoutEOF() {
	LineLength := 0
	for {
		sum := 0
		noContent, _ := fmt.Scan(&LineLength)
		if noContent == 0 {
			break
		} else {
			arr := make([]int, LineLength)
			for i := 0; i < LineLength; i++ {
				fmt.Scan(&arr[i])
			}
			for i := 0; i < LineLength; i++ {
				sum += arr[i]
			}
		}
		fmt.Println(sum)
	}
}

/*
2-2,知道行数，每行数字不固定,但知道数量，结束符
*/

func GetFixedNumberAndLineInputWithLineNumber() {
	LineLength := 0

	fmt.Scan(&LineLength)

	for lineCnt := 0; lineCnt < LineLength; lineCnt++ {
		sum := 0
		eachLineLength := 0
		fmt.Scan(&eachLineLength)
		arr := make([]int, eachLineLength)
		for i := 0; i < eachLineLength; i++ {
			fmt.Scan(&arr[i])
		}
		for i := 0; i < LineLength; i++ {
			sum += arr[i]
		}
		fmt.Println(sum)
	}
}

/*
3-1,整行读，不知道行数，不知道每行元素数量
*/

func RandomReadLine() {
	inputs := bufio.NewScanner(os.Stdin)
	for inputs.Scan() {
		data := strings.Split(inputs.Text(), " ")
		sum := 0
		for _, i := range data {
			val, _ := strconv.Atoi(i)
			sum += val
		}
		fmt.Println(sum)
	}
}

func main() {
	//GetFixedNumberInput()
	//GetFixedNumberAndLineInput()
	RandomReadLine()

	/*
		一道题目有好几个部分，
		1，输入整理
		2，基础函数与工具库，比较大小，字符串处理，字符串转数字
		3，常见算法
	*/
}

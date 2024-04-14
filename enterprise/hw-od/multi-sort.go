package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Height int
	Weight int
	Index  int
}

type StudentList []Student

func (s StudentList) Len() int {
	return len(s)
}

func (s StudentList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s StudentList) Less(i, j int) bool {
	if s[i].Height == s[j].Height {
		return s[i].Weight < s[j].Weight
	}
	return s[i].Height < s[j].Height
}

func main() {
	demo := StudentList{
		{
			Height: 3,
			Weight: 1,
			Index:  1,
		},
		{
			Height: 1,
			Weight: 1,
			Index:  2,
		},
		{
			Height: 3,
			Weight: 2,
			Index:  3,
		},
		{
			Height: 1,
			Weight: 1,
			Index:  4,
		},
	}
	sort.Sort(demo)

	for i := 0; i < demo.Len(); i++ {
		fmt.Printf("%+v\n", demo[i])
	}

}

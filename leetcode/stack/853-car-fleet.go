package main

import "sort"

func carFleet(target int, position []int, speed []int) int {

	list := make([][2]int, len(position))

	for i := 0; i < len(position); i++ {
		list[i][0] = position[i]
		list[i][1] = speed[i]
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i][0] > list[j][0]
	})

	stack := make([]float64, 0)
	for i := 0; i < len(list); i++ {
		needTime := float64(target-list[i][0]) / float64(list[i][1])

		if len(stack) == 0 || needTime > stack[len(stack)-1] {
			stack = append(stack, needTime)
		}
	}

	return len(stack)
}

func carFleetOptimization(target int, position []int, speed []int) int {
	m := make([]int, target+1)
	for i := range position {
		m[position[i]] = speed[i]
	}

	var out int
	var last float64
	for p := target; p >= 0; p-- {
		if m[p] == 0 { //
			continue
		}

		t := float64(target-p) / float64(m[p])
		if last < t {
			last = t
			out++
		}
	}

	return out
}

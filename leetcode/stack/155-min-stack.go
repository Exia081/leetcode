package main

/*
	https://leetcode.com/problems/min-stack/solutions/5028874/very-simple-golang-solution-o-1-time/

思路，同时维护两个栈，一个正常的栈，一个最小栈，最小栈上，每次元素入栈都与栈顶元素比较，选择最小的元素入栈，这样就可以一直记录当前栈中的最小元素的
*/
type MinStack struct {
	stack []int
	min   []int
	top   int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}, -1}
}

func (this *MinStack) Push(val int) {
	if this.top == -1 {
		this.min = append(this.min, val)
	} else {
		this.min = append(this.min, min(val, (this.min)[this.top]))
	}
	this.stack = append(this.stack, val)
	this.top++
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:this.top]
	this.min = this.min[:this.top]
	this.top--
	return
}

func (this *MinStack) Top() int {
	return (this.stack)[this.top]
}

func (this *MinStack) GetMin() int {
	return (this.min)[this.top]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

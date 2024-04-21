package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
1，核心思路，用队列的方式来保存数据
*/

func Bfs(root *TreeNode) (result [][]int) {
	if root == nil {
		return
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	//
	for len(queue) > 0 {

		level := make([]int, 0)

		currentNodeListLength := len(queue)
		for i := 0; i < currentNodeListLength; i++ {

			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}

		result = append(result, level)
	}

	return result
}

func Dfs(root *TreeNode) (result []int) {
	stack := make([]*TreeNode, 0)
	if root != nil {
		stack = append(stack, root)
	}

	result = make([]int, 0)

	for len(stack) > 0 {
		length := len(stack)
		node := stack[length-1]
		stack = stack[:length-1]
		result = append(result, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	return result
}

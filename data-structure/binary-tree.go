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

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func BFS(root *Node) (result [][]int) {
	if root == nil {
		return
	}

	list := make([]*Node, 0)

	list = append(list, root)

	for len(list) > 0 {

		levelList := make([]int, 0)

		levelLength := len(list)

		for i := 0; i < levelLength; i++ {

			levelList = append(levelList, list[i].Val)

			if list[i].Left != nil {
				list = append(list, list[i].Left)
			} else {
				list = append(list, &Node{
					Val: 0,
				})
			}

			if list[i].Right != nil {
				list = append(list, list[i].Right)
			} else {
				list = append(list, &Node{
					Val: 0,
				})
			}
		}

		list = list[levelLength:]

		//判断整个 levelList 是否都是 0

		found := false
		for k := 0; k < len(levelList); k++ {
			if levelList[k] != 0 {
				found = true
			}
		}

		if !found {
			break
		}

		result = append(result, levelList)
	}
	return
}

func CheckValid(result [][]int) bool {

	for level := 0; level < len(result); level++ {

		i, j := 0, len(result[level])-1

		for i <= j {
			if result[level][i] != result[level][j] {
				return false
			}
			i++
			j--
		}
	}
	return true
}

// BuildTree todo 根据前序遍历和中序遍历构建树,有空再对这个补充完成
func BuildTree(preOrder, inOrder []int) (node *Node) {

	if len(inOrder) == 0 {
		return nil
	}

	rootVal := preOrder[0]

	inOrderIndex := -1
	var leftInOrder, rightInOrder []int
	for i := 0; i < len(inOrder); i++ {
		if inOrder[i] == rootVal {
			inOrderIndex = i
			break
		}
	}

	// preorder 1,2,3
	// inorder  2,1,3
	if inOrderIndex != -1 {
		leftInOrder = preOrder[0:inOrderIndex]
		rightInOrder = preOrder[inOrderIndex+1:]
	}

	var leftPreOrder, rightPreOrder []int

	for i := 1; i < len(preOrder); i++ {
		if InArray(preOrder[i], leftInOrder) {
			leftPreOrder = append(leftPreOrder, preOrder[i])
		} else {
			rightPreOrder = append(rightPreOrder, preOrder[i])
		}
	}

	node = &Node{
		Val:   rootVal,
		Left:  BuildTree(leftPreOrder, leftInOrder),
		Right: BuildTree(rightPreOrder, rightInOrder),
	}
	return
}

func InArray(val int, list []int) bool {
	for i := 0; i < len(list); i++ {
		if val == list[i] {
			return true
		}
	}
	return false
}

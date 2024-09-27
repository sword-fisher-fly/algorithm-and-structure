package tree

func InvertTree(root *TreeNode) {
	if root == nil {
		return
	}

	root.Left, root.Right = root.Right, root.Left

	InvertTree(root.Left)
	InvertTree(root.Right)
}

func InvertTreeByDFS(root *TreeNode) {
	if root == nil {
		return
	}

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curNode.Left, curNode.Right = curNode.Right, curNode.Left // root

		if curNode.Right != nil { // right
			stack = append(stack, curNode.Right)
		}

		if curNode.Left != nil { // left
			stack = append(stack, curNode.Left)
		}
	}
}

func InvertTreeByBFS(root *TreeNode) {
	if root == nil {
		return
	}

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		levelSize := len(stack)

		for i := 0; i < levelSize; i++ {
			curNode := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			curNode.Left, curNode.Right = curNode.Right, curNode.Left
			if curNode.Right != nil {
				stack = append(stack, curNode.Right)
			}

			if curNode.Left != nil {
				stack = append(stack, curNode.Left)
			}
		}
	}
}

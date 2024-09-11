package tree

func PostOrderStack(root *TreeNode) []any {
	stackNode1, stackNode2 := []*TreeNode{}, []*TreeNode{}

	if root != nil {
		stackNode1 = append(stackNode1, root)
	}

	for len(stackNode1) != 0 {
		node := stackNode1[len(stackNode1)-1]
		stackNode1 = stackNode1[:len(stackNode1)-1]
		stackNode2 = append(stackNode2, node)

		if node.Left != nil {
			stackNode1 = append(stackNode1, node.Left)
		}

		if node.Right != nil {
			stackNode1 = append(stackNode1, node.Right)
		}
	}

	var ret []any
	for len(stackNode2) > 0 {
		node := stackNode2[len(stackNode2)-1]
		stackNode2 = stackNode2[:len(stackNode2)-1]
		ret = append(ret, node.Val)
	}

	return ret
}

func PostOrderSingleStack(root *TreeNode) []any {
	stack := []*TreeNode{}

	curNode := root
	var preNode *TreeNode

	var ret []any

	for curNode != nil || len(stack) > 0 {
		for curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
		}

		curNode = stack[len(stack)-1]
		if curNode.Right == nil || curNode.Right == preNode {
			stack = stack[:len(stack)-1]
			ret = append(ret, curNode.Val)
			preNode = curNode
			curNode = nil
		} else {
			curNode = curNode.Right
		}
	}

	return ret
}

func PostOrderStackII(root *TreeNode) []any {
	if root == nil {
		return nil
	}

	var ret []any

	stack := []*TreeNode{root}

	for len(stack) != 0 {
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, curNode.Val)

		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
		}

		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
		}
	}

	i := 0
	j := len(ret) - 1
	for i < j {
		ret[i], ret[j] = ret[j], ret[i]
		i++
		j--
	}

	return ret
}

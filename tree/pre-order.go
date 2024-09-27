package tree

// root,left,right
// right clild 位置可以放置在多处
func PreOrderStack(root *TreeNode) []any {
	var ret []any

	curNode := root
	stackNode := make([]*TreeNode, 0)

	for curNode != nil || len(stackNode) > 0 {
		for curNode != nil {
			ret = append(ret, curNode.Val)
			stackNode = append(stackNode, curNode.Right)
			curNode = curNode.Left
		}

		curNode = stackNode[len(stackNode)-1]
		stackNode = stackNode[:len(stackNode)-1]
	}

	return ret
}

func PreOrderStackII(root *TreeNode) []any {
	if root == nil {
		return nil
	}

	var ret []any
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		ret = append(ret, node.Val) //root

		if node.Right != nil {
			stack = append(stack, node.Right) //right
		}

		if node.Left != nil {
			stack = append(stack, node.Left) // left
		}
	}

	return ret
}

// func PreOrderStackIII(root *TreeNode) []any {
// 	var ret []any

// 	curNode := root
// 	stackNode := make([]*TreeNode, 0)

// 	for curNode != nil || len(stackNode) > 0 {
// 		for curNode != nil {
// 			ret = append(ret, curNode.Val)
// 			// if curNode.Right != nil {
// 			// stackNode = append(stackNode, curNode.Right)
// 			// }
// 			curNode = curNode.Left
// 		}

// 		curNode = stackNode[len(stackNode)-1]
// 		stackNode = stackNode[:len(stackNode)-1]

// 		curNode = curNode.Right
// 	}

// 	return ret
// }

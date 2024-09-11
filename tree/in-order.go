package tree

// stack method to in-order traverse
// left, root, right
//
//	 1
//	/ \
//
// 2  3
func InOrderStack(root *TreeNode) []any {
	var ret []any
	curNode := root

	stackNode := make([]*TreeNode, 0)

	for curNode != nil || len(stackNode) != 0 {
		for curNode != nil {
			stackNode = append(stackNode, curNode)
			curNode = curNode.Left
		}

		curNode = stackNode[len(stackNode)-1]
		stackNode = stackNode[:len(stackNode)-1]
		ret = append(ret, curNode.Val)
		curNode = curNode.Right
	}

	return ret
}

func InOrderStackII(root *TreeNode) []any {
	var ret []any

	if root == nil {
		return nil
	}

	stack := []*TreeNode{}
	curNode := root

	for curNode != nil || len(stack) != 0 {
		if curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left //left
		} else {
			curNode = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret = append(ret, curNode.Val) //root
			curNode = curNode.Right        //right
		}
	}

	return ret
}

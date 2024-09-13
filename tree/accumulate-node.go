package tree

// 给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree）
// 使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和

func AccumulateTreeNodeByRecursive(root *TreeNode) *TreeNode {
	preNode := &TreeNode{Val: 0}

	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		traversal(node.Right)
		node.Val = node.Val.(int) + preNode.Val.(int)
		preNode = node
		traversal(node.Left)
	}

	traversal(root)

	return root
}

func AccumulateTreeNodeByIteration(root *TreeNode) *TreeNode {
	preNode := &TreeNode{Val: 0}

	stack := []*TreeNode{}
	curNode := root
	for curNode != nil || len(stack) > 0 {
		if curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Right // 右
		} else {
			curNode = stack[len(stack)-1] // 中
			stack = stack[:len(stack)-1]
			curNode.Val = curNode.Val.(int) + preNode.Val.(int)
			preNode = curNode
			curNode = curNode.Left //左
		}
	}

	return root
}

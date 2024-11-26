package tree

func MergeTrees(root1, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	root1.Val = root1.Val.(int) + root2.Val.(int)
	root1.Left = MergeTrees(root1.Left, root2.Left)
	root1.Right = MergeTrees(root1.Right, root2.Right)

	return root1
}

func MergeTreesByIteration(root1, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	queue := []*TreeNode{root1, root2}

	for len(queue) > 0 {
		curNode1 := queue[0]
		curNode2 := queue[1]
		queue = queue[2:]

		curNode1.Val = curNode1.Val.(int) + curNode2.Val.(int)
		// 2*2 种情况
		if curNode1.Left != nil && curNode2.Left != nil {
			queue = append(queue, curNode1.Left)
			queue = append(queue, curNode2.Left)
		}

		if curNode1.Right != nil && curNode2.Right != nil {
			queue = append(queue, curNode1.Right)
			queue = append(queue, curNode2.Right)
		}

		if curNode1.Left == nil && curNode2.Left != nil {
			curNode1.Left = curNode2.Left
		}

		if curNode1.Right == nil && curNode2.Right != nil {
			curNode1.Right = curNode2.Right
		}
	}

	return root1
}

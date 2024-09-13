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

// TODO:
func MergeTreesByIteration(root1, root2 *TreeNode) *TreeNode {

	return nil
}

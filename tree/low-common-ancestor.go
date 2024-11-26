package tree

// 前提:
// 1) 假设p、q的最近公共祖先存在
// 2) 二叉树为二叉搜索树
func LowestCommonAncestorInBSTByRecursive(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val.(int) < p.Val.(int) && root.Val.(int) < q.Val.(int) {
		return LowestCommonAncestorInBSTByRecursive(root.Right, p, q)
	}

	if root.Val.(int) > p.Val.(int) && root.Val.(int) > q.Val.(int) {
		return LowestCommonAncestorInBSTByRecursive(root.Left, p, q)
	}

	return root
}

// func LowestCommonAncestorInBSTByIteration(root, p, q *TreeNode) *TreeNode {
// 	for root != nil {
// 		if root.Val.(int) < p.Val.(int) && root.Val.(int) < p.Val.(int) {
// 			root = root.Right
// 		} else if root.Val.(int) > p.Val.(int) && root.Val.(int) > q.Val.(int) {
// 			root = root.Left
// 		} else {
// 			return root
// 		}
// 	}

// 	return root
// }

// 普通二叉树, 从底向上回溯
// 前提: 假设p、q为二叉树的节点（有无可能不存在呢？？）
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}

	return left
}

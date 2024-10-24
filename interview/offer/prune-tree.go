package offer

// TODO:
// https://leetcode.cn/problems/binary-tree-pruning/description/
// 注意点
// 1) 必须从底层向上一层层剪枝
func PruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = PruneTree(root.Left)
	root.Right = PruneTree(root.Right)

	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}

	return root
}

package offer

// TODO:
//  1. 查找树最底层最左侧节点
//     https://leetcode.cn/problems/find-bottom-left-tree-value/
func FindBottomLeftValue(root *TreeNode) int {
	return 0
}

// TODO:
//
// https://leetcode.cn/problems/binary-tree-right-side-view/
// 2. 从右侧看二叉树的视图

func RightSideView(root *TreeNode) []int {
	return nil
}

// TODO:
// https://leetcode.cn/problems/find-largest-value-in-each-tree-row/

// 3. 二叉树每层的最大值

func FindLargestValueInEachTreeRow(root *TreeNode) []int {
	return nil
}

// 4. 二叉树根节点到叶子结点的和

func SumNumbers(root *TreeNode) int {

	var dfs func(*TreeNode, int) int
	dfs = func(r *TreeNode, sum int) int {
		if r == nil {
			return 0
		}

		sum = sum*10 + r.Val.(int)
		if r.Left == nil && r.Right == nil {
			return sum
		}

		return dfs(root.Left, sum) + dfs(root.Right, sum)
	}

	return dfs(root, 0)
}

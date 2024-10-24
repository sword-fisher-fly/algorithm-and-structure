package offer

// https://leetcode.cn/problems/path-sum-iii/
func rootSum(r *TreeNode, sum int) int {
	if r == nil {
		return 0
	}

	ret := 0
	if r.Val == sum {
		ret++
	}

	ret += rootSum(r.Left, sum-r.Val.(int))
	ret += rootSum(r.Right, sum-r.Val.(int))

	return ret
}

func PathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	ret := rootSum(root, targetSum)
	ret += PathSum(root.Left, targetSum)
	ret += PathSum(root.Right, targetSum)

	return ret
}

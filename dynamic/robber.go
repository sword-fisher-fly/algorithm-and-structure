package dynamic

import "github.com/sword-fisher-fly/algorithm-and-structure/tree"

// 1）有一排房子，每间房子都有一定现金，如偷相邻两个房屋则触发告警，
// 求在不触动警报的情况下，能够偷窃到的最高金额
// dp[i]: 表示偷盗第i个房屋的最高金额，dp[0]表示第一个房屋
// dp[i] = max(dp[i-1], dp[i-2]+arr[i])
func RobberHouseOnLine(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[0]
	}

	dp := make([]int, len(arr))
	dp[0] = arr[0]
	dp[1] = max(arr[0], arr[1])

	for i := 2; i < len(arr); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+arr[i])
	}

	return dp[len(arr)-1]
}

// 房屋排列为环
// dp[i]: 表示偷盗第i个房屋的最高金额，dp[0]表示第一个房屋
// 1) {arr[0] ... arr[n-2]}
// 2) {arr[1] ... arr[n-1]}
// 3) {arr[1] ... arr[n-2]}
func RobberHoustOnLoop(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[0]
	}

	return max(RobberHouseOnLine(arr[0:len(arr)-1]), RobberHouseOnLine(arr[1:]))
}

// 房屋排列成二叉树
// [3,2,3,-1,3,-1,1] -> 7
// [3,4,5,1,3,-1,1] -> 9
// 1) 偷父节点
// 2）不偷父节点
func RobberHouseOnTree(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return root.Val.(int)
	}
	// 1)偷父节点
	val1 := root.Val.(int)
	if root.Left != nil {
		val1 += RobberHouseOnTree(root.Left.Left) + RobberHouseOnTree(root.Left.Right)
	}

	if root.Right != nil {
		val1 += RobberHouseOnTree(root.Right.Left) + RobberHouseOnTree(root.Right.Right)
	}

	// 2) 不偷父节点
	val2 := RobberHouseOnTree(root.Left) + RobberHouseOnTree(root.Right)
	return max(val1, val2)
}

func robberHouseOnTreeWithMemory(root *tree.TreeNode, umap map[*tree.TreeNode]int) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return root.Val.(int)
	}

	if val, ok := umap[root]; ok {
		return val
	}
	// 1)偷父节点
	val1 := root.Val.(int)
	if root.Left != nil {
		val1 += RobberHouseOnTree(root.Left.Left) + RobberHouseOnTree(root.Left.Right)
	}

	if root.Right != nil {
		val1 += RobberHouseOnTree(root.Right.Left) + RobberHouseOnTree(root.Right.Right)
	}

	// 2)不偷父节点
	val2 := RobberHouseOnTree(root.Left) + RobberHouseOnTree(root.Right)
	umap[root] = max(val1, val2)

	return max(val1, val2)
}
func RobberHouseOnTreeWithMemory(root *tree.TreeNode) int {
	umap := make(map[*tree.TreeNode]int)

	return robberHouseOnTreeWithMemory(root, umap)
}

func robberHouseOnTreeByDP(root *tree.TreeNode) [2]int {
	if root == nil {
		return [2]int{0, 0}
	}
	left := robberHouseOnTreeByDP(root.Left)
	right := robberHouseOnTreeByDP(root.Right)

	// 1) 偷父节点
	val1 := root.Val.(int) + left[1] + right[1]
	// 2) 不偷父节点
	val2 := max(left[0], left[1]) + max(right[0], right[1])

	return [2]int{val1, val2}
}

// dp[0], dp[1]
// dp[0]: 偷该该节点
// dp[1]: 不偷该节点
func RobberHouseOnTreeByDP(root *tree.TreeNode) int {
	result := robberHouseOnTreeByDP(root)

	return max(result[0], result[1])
}

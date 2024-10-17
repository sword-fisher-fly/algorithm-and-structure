package recursive

type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

type BSTree struct {
	root *TreeNode
}

// 1) 有序数组中间值作为root节点
// 2) 以中间值索引切割为左右两个子树
// 3) 递归构建左右子树

func ConstructBSTreeFromSortedArray(nums []int) *BSTree {
	// [start, end]
	var buildBSTree func(nums []int, start, end int) *TreeNode

	buildBSTree = func(nums []int, start, end int) *TreeNode {
		if start > end {
			return nil
		}

		mid := (start + end) / 2
		root := &TreeNode{Val: nums[mid]}
		root.Left = buildBSTree(nums, start, mid-1)
		root.Right = buildBSTree(nums, mid+1, end)

		return root
	}

	root := buildBSTree(nums, 0, len(nums)-1)

	return &BSTree{root: root}
}

func InOrderTraverse(root *TreeNode) []interface{} {
	var ans []interface{}

	var inOrder func(root *TreeNode)

	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inOrder(root.Left)
		ans = append(ans, root.Val)
		inOrder(root.Right)
	}

	inOrder(root)

	return ans
}

func LevelTraversal(root *TreeNode) [][]interface{} {
	var ans [][]interface{}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		level := []interface{}{}
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, level)
	}

	return ans
}

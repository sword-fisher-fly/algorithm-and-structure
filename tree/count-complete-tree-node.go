package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	level := 0
	node := root
	for node.Left != nil {
		level++
		node = node.Left
	}

	exists := func(t *TreeNode, lev int, k int) bool {
		bits := 1 << (lev - 1)
		node := t
		for node != nil && bits > 0 {
			if bits&k == 0 {
				node = node.Left
			} else {
				node = node.Right
			}

			bits >>= 1
		}

		return node != nil
	}

	low := 1 << level
	high := 1<<(level+1) - 1
	// fmt.Printf("level=%d, low=%d, high=%d\n", level, low, high)
	for low < high {
		mid := low + (high-low+1)>>1
		if exists(root, level, mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}

	return low
}

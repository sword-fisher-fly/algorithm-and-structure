package tree

import (
	"fmt"
	"math"
)

// 根节点深度约定为0
func FindLeftBottomLeafByRecursive(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	maxDepth := math.MinInt
	var result *TreeNode

	var traversal func(node *TreeNode, depth int)
	traversal = func(node *TreeNode, depth int) {
		if node != nil && node.Left == nil && node.Right == nil {
			if depth > maxDepth {
				maxDepth = depth
				result = node
			}

			return
		}

		// 先遍历左孩子节点
		if node.Left != nil {
			traversal(node.Left, depth+1)
		}

		if node.Right != nil {
			traversal(node.Right, depth+1)
		}
	}

	traversal(root, 0)

	fmt.Printf("maxDepth: %d\n", maxDepth)

	return result
}

func FindLeftBottomLeafByIteration(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var result *TreeNode

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			curNode := queue[0]
			queue = queue[1:]

			if i == 0 {
				result = curNode
			}

			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}

			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}
	}

	return result
}

package tree

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

		return dfs(r.Left, sum) + dfs(r.Right, sum)
	}

	return dfs(root, 0)
}

func SumNumbersByLevelTraversal(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	type pair struct {
		node    *TreeNode
		pathSum int
	}

	queue := []pair{{node: root, pathSum: root.Val.(int)}}

	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]

		if curNode.node.Left == nil && curNode.node.Right == nil {
			sum += curNode.pathSum
		}

		if curNode.node.Left != nil {
			queue = append(queue, pair{node: curNode.node.Left, pathSum: curNode.pathSum*10 + curNode.node.Left.Val.(int)})
		}

		if curNode.node.Right != nil {
			queue = append(queue, pair{node: curNode.node.Right, pathSum: curNode.pathSum*10 + curNode.node.Right.Val.(int)})
		}
	}

	return sum
}

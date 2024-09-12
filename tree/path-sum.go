package tree

func HasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	var traversal func(node *TreeNode, curSum int) bool

	traversal = func(node *TreeNode, count int) bool {
		if node.Left == nil && node.Right == nil && count == 0 {
			return true
		}

		if node.Left == nil && node.Right == nil {
			return false
		}

		if node.Left != nil {
			count -= node.Left.Val.(int)
			if traversal(node.Left, count) {
				return true
			}

			count += node.Left.Val.(int)
		}

		if node.Right != nil {
			count -= node.Right.Val.(int)
			if traversal(node.Right, count) {
				return true
			}

			count += node.Right.Val.(int)
		}

		return false
	}

	return traversal(root, sum-root.Val.(int))
}

func HasPathSumSimple(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && root.Val.(int) == sum {
		return true
	}

	return HasPathSumSimple(root.Left, sum-root.Val.(int)) || HasPathSumSimple(root.Right, sum-root.Val.(int))
}

func HasPathSumByIteration(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	type Item struct {
		node    *TreeNode
		pathSum int
	}

	stack := []Item{{root, root.Val.(int)}}

	for len(stack) != 0 {
		curItem := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// fmt.Printf("curItem: %v, pathSum: %v\n", curItem.node.Val, curItem.pathSum)
		if curItem.node.Left == nil && curItem.node.Right == nil && curItem.pathSum == sum {
			return true
		}

		if curItem.node.Right != nil {
			stack = append(stack, Item{curItem.node.Right, curItem.pathSum + curItem.node.Right.Val.(int)})
		}

		if curItem.node.Left != nil {
			stack = append(stack, Item{curItem.node.Left, curItem.pathSum + curItem.node.Left.Val.(int)})
		}
	}

	return false
}

func PathSum(root *TreeNode, sum int) [][]any {
	result := make([][]any, 0)
	if root == nil {
		return result
	}

	var traversal func(node *TreeNode, curSum int, path []any)

	traversal = func(node *TreeNode, count int, path []any) {
		if node.Left == nil && node.Right == nil && count == 0 {
			temp := make([]any, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		if node.Left != nil {
			path = append(path, node.Left.Val)
			traversal(node.Left, count-node.Left.Val.(int), path)
			path = path[:len(path)-1]
		}

		if node.Right != nil {
			path = append(path, node.Right.Val)
			traversal(node.Right, count-node.Right.Val.(int), path)
			path = path[:len(path)-1]
		}
	}

	traversal(root, sum-root.Val.(int), []any{root.Val})

	return result
}

// TODO: 迭代方式记录所有路径？？

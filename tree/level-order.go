package tree

func (t *Tree) LevelTraversal() [][]any {
	return LevelTraversal(t.root)
}
func LevelTraversal(root *TreeNode) [][]any {
	if root == nil {
		return nil
	}

	stack := []*TreeNode{root}

	ret := [][]any{}
	level := 0
	for len(stack) != 0 {
		size := len(stack)

		ret = append(ret, []any{})
		for i := 0; i < size; i++ {
			curNode := stack[0]
			stack = stack[1:]
			ret[level] = append(ret[level], curNode.Val)

			if curNode.Left != nil {
				stack = append(stack, curNode.Left)
			}

			if curNode.Right != nil {
				stack = append(stack, curNode.Right)
			}
		}

		level++
	}

	return ret
}

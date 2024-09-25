package tree

func serializeTreeNodeList(bi *TreeNode) []any {
	if bi == nil {
		return []any{}
	}

	res := []any{}

	for bi != nil {
		res = append(res, bi.Val)
		bi = bi.Right
	}

	return res
}
func Convert2BiNode(t *TreeNode) *TreeNode {
	if t == nil {
		return nil
	}

	dummyNode := &TreeNode{Right: t}
	preNode := dummyNode
	// fmt.Printf("Initialize preNode=%v\n", preNode.Val)

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		// fmt.Printf("preNode=%v\n", preNode.Val)
		preNode.Right = root
		root.Left = nil
		preNode = root
		dfs(root.Right)
	}

	// dfs(root) //ok
	dfs(dummyNode.Right)

	return dummyNode.Right
}

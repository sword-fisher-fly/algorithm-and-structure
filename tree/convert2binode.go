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
func Convert2BiNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	dummyNode := &TreeNode{Right: root}
	preNode := dummyNode

	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		preNode.Right = root
		root.Left = nil
		preNode = root
		dfs(root.Right)
	}

	// dfs(root) //ok
	dfs(dummyNode.Right)

	return dummyNode.Right
}

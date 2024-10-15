package tree

func IsSame(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && IsSame(p.Left, q.Left) && IsSame(p.Right, q.Right)
}

func IsSameByBFS(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	queue := []*TreeNode{p, q}
	for len(queue) > 0 {
		t1 := queue[0]
		t2 := queue[1]
		queue = queue[2:]

		if t1 == nil && t2 == nil {
			continue
		}

		if t1 == nil || t2 == nil || t1.Val != t2.Val {
			return false
		}

		queue = append(queue, t1.Left, t2.Left)
		queue = append(queue, t1.Right, t2.Right)
	}

	return true
}

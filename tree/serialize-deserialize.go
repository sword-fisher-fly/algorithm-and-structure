package tree

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/solutions/290065/er-cha-shu-de-xu-lie-hua-yu-fan-xu-lie-hua-by-le-2/
// type Codec struct{}

// func Constructor() (_ Codec) {
// 	return
// }

// func (c Codec) serialize(root *TreeNode) string {
// 	if root == nil {
// 		return "X"
// 	}
// 	left := "(" + c.serialize(root.Left) + ")"
// 	right := "(" + c.serialize(root.Right) + ")"
// 	return left + strconv.Itoa(root.Val.(int)) + right
// }

// func (Codec) deserialize(data string) *TreeNode {
// 	var parse func() *TreeNode
// 	parse = func() *TreeNode {
// 		if data[0] == 'X' {
// 			data = data[1:]
// 			return nil
// 		}
// 		node := &TreeNode{}
// 		data = data[1:] // 跳过左括号
// 		node.Left = parse()
// 		data = data[1:] // 跳过右括号
// 		i := 0
// 		for data[i] == '-' || '0' <= data[i] && data[i] <= '9' {
// 			i++
// 		}
// 		node.Val, _ = strconv.Atoi(data[:i])
// 		data = data[i:]
// 		data = data[1:] // 跳过左括号
// 		node.Right = parse()
// 		data = data[1:] // 跳过右括号
// 		return node
// 	}

// 	return parse()
// }

type Codec struct{}

func Constructor() (_ Codec) {
	return
}

func (Codec) serialize(root *TreeNode) string {
	sb := &strings.Builder{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(strconv.Itoa(node.Val.(int)))
		sb.WriteByte(',')
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	ret := sb.String()

	return ret[:len(ret)-1]
}

func (Codec) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil // 必然经过此处
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &TreeNode{val, build(), build()}
	}

	return build()
}

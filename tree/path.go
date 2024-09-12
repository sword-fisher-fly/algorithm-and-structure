package tree

import (
	"fmt"
	"strconv"
)

func toStr(arr []any) []string {
	var ret []string
	for _, v := range arr {
		ret = append(ret, fmt.Sprintf("%v", v))
	}

	return ret
}

// 错误样板，不知道原因??
// func treePathFromRootToLeafWithError(root *TreeNode, path []any, result *[]string) {
// 	if root.Left == nil && root.Right == nil {
// 		*result = append(*result, strings.Join(toStr(path), "->"))
// 		return
// 	}

// 	path = append(path, root.Val)

// 	if root.Left != nil {
// 		treePathFromRootToLeafWithError(root.Left, path, result)
// 		path = path[:len(path)-1]
// 	}

//		if root.Right != nil {
//			treePathFromRootToLeafWithError(root.Right, path, result)
//			path = path[:len(path)-1]
//		}
//	}
func treePathFromRootToLeaf(root *TreeNode, path string, result *[]string) {
	if root.Left == nil && root.Right == nil {
		*result = append(*result, path+strconv.Itoa(root.Val.(int)))
		return
	}

	path += strconv.Itoa(root.Val.(int))

	if root.Left != nil {
		treePathFromRootToLeaf(root.Left, path+"->", result)
	}

	if root.Right != nil {
		treePathFromRootToLeaf(root.Right, path+"->", result)
	}
}

func TreePathFromRootToLeaf(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	result := []string{}
	path := ""

	treePathFromRootToLeaf(root, path, &result)

	return result
}

// copy
func TreePathFromRootToLeafII(root *TreeNode) []string {
	res := make([]string, 0)
	var travel func(node *TreeNode, s string)
	travel = func(node *TreeNode, s string) {
		if node.Left == nil && node.Right == nil {
			v := s + strconv.Itoa(node.Val.(int))
			res = append(res, v)
			return
		}
		s = s + strconv.Itoa(node.Val.(int)) + "->"
		if node.Left != nil {
			travel(node.Left, s)
		}
		if node.Right != nil {
			travel(node.Right, s)
		}
	}

	travel(root, "")
	return res
}

func TreePathFromRootToLeafByIteration(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	treeStack := []*TreeNode{root}
	pathStack := []string{strconv.Itoa(root.Val.(int))}

	res := []string{}
	for len(treeStack) != 0 {
		curNode := treeStack[len(treeStack)-1]
		treeStack = treeStack[:len(treeStack)-1]
		path := pathStack[len(pathStack)-1]
		pathStack = pathStack[:len(pathStack)-1]
		if curNode.Left == nil && curNode.Right == nil {
			res = append(res, path)
			continue
		}

		if curNode.Right != nil {
			treeStack = append(treeStack, curNode.Right)
			pathStack = append(pathStack, path+"->"+strconv.Itoa(curNode.Right.Val.(int)))
		}

		if curNode.Left != nil {
			treeStack = append(treeStack, curNode.Left)
			pathStack = append(pathStack, path+"->"+strconv.Itoa(curNode.Left.Val.(int)))
		}
	}

	return res
}

package tree

// https://www.programmercarl.com/0110.%E5%B9%B3%E8%A1%A1%E4%BA%8C%E5%8F%89%E6%A0%91.html#%E7%AE%97%E6%B3%95%E5%85%AC%E5%BC%80%E8%AF%BE

// ?? 有问题
// func GetDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}

// 	stack := []*TreeNode{root}
// 	depth, result := 0, 0

// 	for len(stack) != 0 {
// 		curNode := stack[len(stack)-1]
// 		if curNode != nil {
// 			stack := stack[:len(stack)-1]
// 			stack = append(stack, curNode)
// 			stack = append(stack, nil)
// 			depth++

// 			if curNode.Right != nil {
// 				stack = append(stack, curNode.Right)
// 			}

// 			if curNode.Left != nil {
// 				stack = append(stack, curNode.Left)
// 			}
// 		} else {
// 			stack = stack[:len(stack)-1]
// 			curNode = stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]
// 			depth--
// 		}

// 		if result < depth {
// 			result = depth
// 		}
// 	}

// 	return result
// }

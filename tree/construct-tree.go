package tree

//	  1
//	 / \
//	2   3
//	/\  /\
//
// 4 5  6 7
// pre-order: root -> left -> right [1,2,4,5,3,6,7]
// in-order: left -> root -> right [4,2,5,1,6,3,7]
// post-order: left -> right -> root [4,5,2,6,7,3,1]
//  1. root = 1  pre-order分为[2,4,5] [3,6,7] in-order分为[4,2,5] [6,3,7]
//  2. 两组 pre-order: [2,4,5] in-order: [4,2,5] pre-order: [3,6,7] in-order: [6,3,7]
//     root = 2 pre-order分为两组 [4] [5] in-order分为两组 [4] [5]
//     root = 3 pre-order分为两组 [6] [7] in-order分为两组 [6] [7]
func ConstructTreeFromPreOrderAndInOrder(preorder []int, inorder []int) *TreeNode {
	var constructTreeFromPreOrderAndInOrder func(preorder []int, preStart int, preEnd int, inorder []int, inStart int, inEnd int) *TreeNode

	constructTreeFromPreOrderAndInOrder = func(preorder []int, preStart int, preEnd int, inorder []int, inStart int, inEnd int) *TreeNode {
		//fmt.Printf("preStart=%d, preEnd=%d, inStart=%d, inEnd=%d\n", preStart, preEnd, inStart, inEnd)
		root := &TreeNode{Val: preorder[preStart]}
		if preEnd-preStart == 1 {
			return root
		}

		delemiterIndex := 0
		for i := inStart; i < inEnd; i++ {
			if preorder[preStart] == inorder[i] {
				delemiterIndex = i
			}
		}

		//fmt.Printf("delemiterIndex=%d, preStart+1=%d, preStart+1+(delemiterIndex-inStart)=%d\n", delemiterIndex, preStart+1, preStart+1+(delemiterIndex-inStart))
		root.Left = constructTreeFromPreOrderAndInOrder(preorder, preStart+1, preStart+1+(delemiterIndex-inStart), inorder, inStart, delemiterIndex)
		root.Right = constructTreeFromPreOrderAndInOrder(preorder, preStart+1+(delemiterIndex-inStart), preEnd, inorder, delemiterIndex+1, inEnd)

		return root
	}

	return constructTreeFromPreOrderAndInOrder(preorder, 0, len(preorder), inorder, 0, len(inorder))
}

// in-order: left -> root -> right [4,2,5,1,6,3,7]
// post-order: left -> right -> root [4,5,2,6,7,3,1]
func ConstructTreeFromInOrderAndPostOrder(inorder []int, postorder []int) *TreeNode {
	var constructTreeFromInOrderAndPostOrder func(inorder []int, inStart int, inEnd int, postorder []int, postStart int, postEnd int) *TreeNode

	constructTreeFromInOrderAndPostOrder = func(inorder []int, inStart int, inEnd int, postorder []int, postStart int, postEnd int) *TreeNode {
		// fmt.Printf("inStart=%d, inEnd=%d, postStart=%d, postEnd=%d\n", inStart, inEnd, postStart, postEnd)
		root := &TreeNode{Val: postorder[postEnd-1]}
		if postEnd-postStart == 1 {
			return root
		}

		delemiterIndex := 0
		for i := inStart; i < inEnd; i++ {
			if postorder[postEnd-1] == inorder[i] {
				delemiterIndex = i
			}
		}

		root.Left = constructTreeFromInOrderAndPostOrder(inorder, inStart, delemiterIndex, postorder, postStart, postStart+(delemiterIndex-inStart))
		root.Right = constructTreeFromInOrderAndPostOrder(inorder, delemiterIndex+1, inEnd, postorder, postStart+(delemiterIndex-inStart), postEnd-1)

		return root
	}

	return constructTreeFromInOrderAndPostOrder(inorder, 0, len(inorder), postorder, 0, len(postorder))
}

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
		if preStart > preEnd || inStart > inEnd {
			return nil
		}

		root := &TreeNode{Val: preorder[preStart]}

		delimeterIdx := 0
		for i := inStart; i <= inEnd; i++ {
			if preorder[preStart] == inorder[i] {
				delimeterIdx = i
				break
			}
		}

		leftSize := delimeterIdx - inStart
		rightSize := inEnd - delimeterIdx
		//fmt.Printf("delimeterIdx=%d, preStart+1=%d, preStart+1+(delimeterIdx-inStart)=%d\n", delimeterIdx, preStart+1, preStart+1+(delimeterIdx-inStart))
		// root.Left = constructTreeFromPreOrderAndInOrder(preorder, preStart+1, preStart+1+(delimeterIdx-inStart), inorder, inStart, delimeterIdx)
		// root.Right = constructTreeFromPreOrderAndInOrder(preorder, preStart+1+(delimeterIdx-inStart), preEnd, inorder, delimeterIdx+1, inEnd)

		root.Left = constructTreeFromPreOrderAndInOrder(preorder, preStart+1, preStart+leftSize, inorder, inStart, inStart+leftSize-1)
		// root.Right = constructTreeFromPreOrderAndInOrder(preorder, preEnd-rightSize+1, preEnd, inorder, delimeterIdx+1, inEnd) // ok ??
		root.Right = constructTreeFromPreOrderAndInOrder(preorder, preEnd-rightSize+1, preEnd, inorder, inEnd-rightSize+1, inEnd)
		return root
	}

	return constructTreeFromPreOrderAndInOrder(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
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

		delimeterIdx := 0
		for i := inStart; i < inEnd; i++ {
			if postorder[postEnd-1] == inorder[i] {
				delimeterIdx = i
			}
		}

		root.Left = constructTreeFromInOrderAndPostOrder(inorder, inStart, delimeterIdx, postorder, postStart, postStart+(delimeterIdx-inStart))
		root.Right = constructTreeFromInOrderAndPostOrder(inorder, delimeterIdx+1, inEnd, postorder, postStart+(delimeterIdx-inStart), postEnd-1)

		return root
	}

	return constructTreeFromInOrderAndPostOrder(inorder, 0, len(inorder), postorder, 0, len(postorder))
}

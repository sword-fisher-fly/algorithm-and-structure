package tree

var (
	level1Arr = []any{1}
	level2Arr = []any{1, 2, 3}
	level3Arr = []any{1, 2, 3, 4, 5, 6, 7}
	level4Arr = []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	level4SymmetricArr = []any{1, 2, 2, 3, 4, 4, 3, 5, -1, -1, -1, -1, -1, -1, 5}
	level4RandomArr    = []any{1, 2, 6, 3, 4, -1, -1, 5}
	//    1
	//   / \
	//   2 6
	//  /\
	// 3 4
	// /
	// 5
	level1PreOrder = []any{1}
	level2PreOrder = []any{1, 2, 3}
	level3PreOrder = []any{1, 2, 4, 5, 3, 6, 7}
	level4PreOrder = []any{1, 2, 4, 8, 9, 5, 10, 11, 3, 6, 12, 13, 7, 14, 15}

	//      1
	//     / \
	//    2  2
	//   /\ /\
	//  3 4 4  3
	//  /       \
	//  5       5
	//         1
	//      /      \
	//     2        3
	//   /  \      /   \
	//  4    5    6     7
	//  /\  / \   /\    /\
	// 8 9 10 11 12 13 14 15
	// http://www.easycode.top/
	// preOrder:  ABDHIEJCFKG
	// inOrder:   HDIBEJAFKCG
	// postOrder: HIDJEBKFGCA
	alphabetArr = []any{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', -1, 'J', -1, 'K'}
)

var (
	level1FullTree      *Tree
	level2FullTree      *Tree
	level3FullTree      *Tree
	level4FullTree      *Tree
	level4RandomTree    *Tree
	level4SymmetricTree *Tree
)

func init() {
	level1FullTree = NewTreeFromArray(level1Arr)
	level2FullTree = NewTreeFromArray(level2Arr)
	level3FullTree = NewTreeFromArray(level3Arr)
	level4FullTree = NewTreeFromArray(level4Arr)

	level4RandomTree = NewTreeFromArray(level4RandomArr)
	level4SymmetricTree = NewTreeFromArray(level4SymmetricArr)
}

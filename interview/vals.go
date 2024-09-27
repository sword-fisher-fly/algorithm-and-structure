package interview

var (
	// 左 -> 下 -> 右 -> 上
	// (x,y)
	directions = [4][2]int{
		{1, 0},
		{0, -1},
		{-1, 0},
		{0, 1},
	}

	directionsII = []int{1, 0, -1, 0, 1}
	// for i := 1; i < len(directionsII); i++ {
	//     (x, y) = (directionsII[i-1], directionsII[i])
	// }
)

var (
	RedColor   = "\033[31m"
	ResetColor = "\033[0m"
)

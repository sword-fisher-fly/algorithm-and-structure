package math

// 一个数组包含红色、白色、蓝色三种颜色，按照红色、白色、蓝色的顺序进行排序。
// red: 0
// white: 1
// blue: 2
// Input: [2,0,2,1,1,0]
// Output: [0,0,1,1,2,2]

// 巧妙的解法: 数组元素值只有3个
func SortColors(colors []int) []int {
	zero, one := 0, 0

	for i, v := range colors {
		colors[i] = 2
		if v <= 1 {
			colors[one] = 1
			one++
		}

		if v == 0 {
			colors[zero] = 0
			zero++
		}
	}

	return colors
}

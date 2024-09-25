package greedy

// 给定一个非负整数数组, 最初位于数组的第一个位置
// 数组中的每个元素代表你在该位置可以跳跃的最大长度

// 分析:
//
//	迭代寻找最大可以覆盖的最大范围, 如最大范围可达数组最后一个元素则返回true
func CanJump(nums []int) bool {
	cover := 0
	if len(nums) < 2 {
		return true
	}

	for i := 0; i <= cover; i++ {
		cover = max(cover, i+nums[i])
		if cover >= len(nums)-1 {
			return true
		}
	}

	return false
}

// 给定一个非负整数数组,最初位于数组第一个位置
// 数组每个元素代表在该位置能跳跃的最大长度
// 求解: 使用最少的跳跃次数达到数组的最后一个位置

// [2, 3, 1, 1, 4]
func JumpMinSteps(nums []int) int {
	curDistance, nextDistance := 0, 0

	var ans int
	for i := 0; i < len(nums); i++ {
		nextDistance = max(nextDistance, nums[i]+i)
		if i == curDistance {
			// fmt.Printf("i=%d, num[%d]=%d, curDistance=%d, nextDistance=%d\n", i, i, nums[i], curDistance, nextDistance)
			curDistance = nextDistance
			ans++
			if curDistance >= len(nums)-1 {
				break
			}
		}
	}

	return ans
}

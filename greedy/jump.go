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

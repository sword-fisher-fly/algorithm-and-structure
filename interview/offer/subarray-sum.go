package offer

// 给定数组nums和目标和, 求有多少连续的子数组和为目标值

// Sum(nums[j..i]) = target
// 1) 暴力求解 直观
// 倒序计算以nums[i]为结尾子数组和为target的数量
func SubArraySum(nums []int, target int) int {
	count := 0
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum = 0
		for end := i; end >= 0; end-- {
			sum += nums[end]
			if sum == target {
				count++
			}
		}
	}

	return count
}

// 哈希表+前缀和
// pre[i] = pre[i-1] + nums[i]
// Sum(nums[j..i]) = pre[i] - pre[j-1] = k
// pre[j-1] = pre[i]-k condition j < i
// mp 记录mp[pre[i]-k]的数量 妙处在这里

func SunArraySumWithHashPrefixSum(nums []int, target int) int {
	mp := map[int]int{0: 1}

	count := 0
	pre := 0
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		// if v, ok := mp[pre-target]; ok {
		// 	count += v
		// }
		count += mp[pre-target]

		mp[pre]++
	}

	return count
}

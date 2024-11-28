package hash

// func subarraySum(nums []int, k int) int {
//     res := 0
//     // sum := 0
//     var backtracing func(int,int)
//     backtracing = func(startIdx int, sum int) {
//         if sum == k {
//            res++
//         }

//         for i := startIdx; i < len(nums); i++ {
//             sum += nums[i]
//             backtracing(i+1, sum)
//             sum -= nums[i]
//         }
//     }

//     backtracing(0, 0)

//     return res
// }

//https://leetcode.cn/problems/subarray-sum-equals-k/?envType=problem-list-v2&envId=2cktkvj
func subarraySum(nums []int, k int) int {
	mp := make(map[int]int)
	preSum := 0
	count := 0
	mp[0] = 1
	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		count += mp[preSum-k]
		mp[preSum] += 1
	}

	return count
}

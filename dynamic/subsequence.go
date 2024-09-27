package dynamic

// 1)  最长上升子序列
// 2） 最长连续递增子序列
// 3） 最长重复子数组
// 4） 最长公共子序列
// 5） 最长子序和
// 6） 判断子序列
// 7） 不同的子序列
// 8)  不相交的线
// 9） 最长回文子序列

// 1) 最大上升递增子序列
// dp[i]: 以i结尾的最长上升子序列的长度

// The same with `LongestAscendingSubSequence` implemented in longest-ascending-sub-sequence.go
func MaxLengthOfIncreasingSubSequence(nums []int) int {
	result := 1

	dp := make([]int, len(nums))
	dp[0] = 1

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(result, dp[i])
	}

	return result
}

// 3） 最长重复子数组的长度
//  nums1: [1,2,3,2,1]
//  nums2: [3,2,1,4,7]
//  Output：[3,2,1]

// dp[i][j]: nums1[:i]和nums2[:j]的最长重复子数组的长度
// the same with `MaxSubPalindrome` defined in max-common-sub-array.go
func MaxRepeatedSubArray(nums1, nums2 []int) int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}

	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	result := 0
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			// } else {
			// 	dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			// }
			if dp[i][j] > result {
				result = dp[i][j]
			}
		}
	}

	return result
}

func MaxRepeatedSubArrayII(nums1, nums2 []int) int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}

	dp := make([]int, len(nums2)+1)

	result := 0
	for i := 1; i <= len(nums1); i++ {
		for j := len(nums2); j > 0; j-- {
			if nums1[i-1] == nums2[j-1] {
				dp[j] = dp[j-1] + 1
			} else {
				dp[j] = 0
			}

			if dp[j] > result {
				result = dp[j]
			}
		}
	}

	return result
}

// 4） 最长公共子序列
// dp[i][j]: nums1[:i]和nums2[:j]的最长公共子序列的长度
func LongestCommonSubSequence(nums1, nums2 []int) int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}

	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(nums1)][len(nums2)]
}

func LongestCommonSubSequenceII(s1 string, s2 string) int {
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}

	dp := make([][]int, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(s1)][len(s2)]
}

// 5） 最长子序和
// dp[i]: 以nums[i-1]结尾的最长子序列和
func MaxSumOfSubSequence(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]

	result := dp[0]
	for i := 1; i < len(nums); i++ {
		// 适合最长上升子序列？？***子序列和是连续的***
		// 上升子序列可以非连续
		// for j := 0; j < i; j++ {
		// 	if dp[j] > 0 {
		// 		dp[i] = dp[j] + nums[i]
		// 	} else {
		// 		dp[i] = nums[i]
		// 	}

		// 	if dp[i] > result {
		// 		result = dp[i]
		// 	}
		// }

		dp[i] = max(dp[i-1]+nums[i], nums[i])
		result = max(result, dp[i])
	}

	return result
}

// 6） 判断子序列
// 判断s是否是t的子序列字符串
// dp[i][j]: s[:i]和t[:j]最大相同子序列的长度
func IsSubsequence(s, t string) bool {
	if len(s) == 0 || len(t) == 0 {
		return false
	}

	if len(s) > len(t) {
		return false
	}

	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 剔除t[j-1], s是不变的
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[len(s)][len(t)] == len(s)
}

// TODO:
// 7) 不相同子序列
// https://www.programmercarl.com/0115.%E4%B8%8D%E5%90%8C%E7%9A%84%E5%AD%90%E5%BA%8F%E5%88%97.html#%E6%80%9D%E8%B7%AF
// 字符串s在字符串t子序列中出现的次数
// dp[i][j]: s[:i]中出现t[:j]形式子序列次数
func NumOfSubsequence(s, t string) int {
	if len(s) == 0 {
		return 0
	}

	if len(t) == 0 {
		return 1
	}

	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}

	for i := 0; i <= len(s); i++ {
		dp[i][0] = 1
	}

	for j := 1; j <= len(t); j++ {
		dp[0][j] = 0
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	// fmt.Printf("s=%s, t=%s, dp=%v\n", s, t, dp)

	return dp[len(s)][len(t)]
}

// 8)  不相交的线
// 与最长公共子序列相同

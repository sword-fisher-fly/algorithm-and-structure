package special

import (
	"strconv"
)

func DecodeNumber(nums string) int {
	// write code here
	if len(nums) < 2 {
		if len(nums) == 1 && nums[0] == '0' {
			return 0
		}

		return len(nums)
	}

	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	val, _ := strconv.Atoi(nums[:2])
	if val == 10 || val >= 27 && val <= 99 {
		dp[1] = 1
	}

	if val >= 11 && val <= 26 {
		dp[1] = 2
	}

	for i := 2; i < n; i++ {
		if nums[i] == '0' {
			if nums[i-1] != '1' && nums[i-1] != '2' {
				return 0
			}

			dp[i] = dp[i-2]
		} else if nums[i] >= '1' && nums[i] <= '6' {
			if nums[i-1] >= '1' && nums[i-1] <= '2' {
				dp[i] = dp[i-2] + dp[i-1]
			} else {
				dp[i] = dp[i-1]
			}
		} else {
			if nums[i-1] == '1' {
				dp[i] = dp[i-2] + dp[i-1]
			} else {
				dp[i] = dp[i-1]
			}
		}
	}

	return dp[n-1]
}

func DecodeNumberII(nums string) int {
	if nums == "0" {
		return 0
	}

	if nums == "10" || nums == "20" {
		return 1
	}

	// 1) exclude the string that can not be decoded
	for i := 1; i < len(nums); i++ {
		if nums[i] == '0' {
			if nums[i-1] != '1' && nums[i-1] != '2' {
				return 0
			}
		}
	}

	dp := make([]int, len(nums)+1)
	dp[0], dp[1] = 1, 1
	// 1) consider the special case first.
	for i := 2; i <= len(nums); i++ {
		if nums[i-2] == '1' && nums[i-1] != '0' || nums[i-2] == '2' && nums[i-1] > '0' && nums[i-1] < '7' {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[len(nums)]
}

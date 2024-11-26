package special

//https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/)

// You have d dice, and each die has f faces numbered 1, 2, ..., f.

// Return the number of possible ways (out of fd total ways) modulo 10^9 + 7 to roll the dice so the sum of the face up numbers equals target.

// Example 1:

// ```text
// Input: d = 1, f = 6, target = 3
// Output: 1
// Explanation:
// You throw one die with 6 faces.  There is only one way to get a sum of 3.
// ```

const mod = 1e9 + 7

func NumRollsToTarget(dices, faces, target int) int {
	dp := [31][1001]int{}
	dp[0][0] = 1

	for d := 1; d <= dices; d++ {
		maxT := min(target, d*faces)
		for t := d; t <= maxT; t++ {
			maxF := min(faces, t)
			for f := 1; f <= maxF; f++ {
				dp[d][t] += dp[d-1][t-f]
			}
			dp[d][t] %= mod
		}
	}

	return dp[dices][target]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

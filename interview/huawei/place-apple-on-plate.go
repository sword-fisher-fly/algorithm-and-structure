package huawei

//https://www.nowcoder.com/practice/bfd8234bb5e84be0b493656e390bdebf?tpId=37&tqId=21284&rp=1&ru=/exam/oj/ta&qru=/exam/oj/ta&sourceUrl=%2Fexam%2Foj%2Fta%3Fpage%3D2%26tpId%3D37%26type%3D37&difficulty=undefined&judgeStatus=undefined&tags=&title=
// m: 物品总数
// n: 盘子数量
func PlaceAppleOnPlateByRecursive(m, n int) int {
	if m == 0 || m == 1 || n == 1 {
		return 1
	}

	if n > m {
		return PlaceAppleOnPlateByRecursive(m, m)
	}

	return PlaceAppleOnPlateByRecursive(m, n-1) + PlaceAppleOnPlateByRecursive(m-n, n)
}

func PlaceAppleOnPlateByDP(m, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// 一个盘子
	for i := 0; i <= m; i++ {
		dp[i][1] = 1
	}

	// 0个或1个苹果
	for i := 1; i <= n; i++ {
		dp[0][i] = 1
		dp[1][i] = 1
	}
	// m个盘子, n个苹果
	// dp[i][j]: i个苹果, j个盘子
	// dp[i][j] = dp[i][i] i < j
	// dp[i][j] = dp[i][j-1]+dp[i-j][j]
	for i := 2; i <= m; i++ {
		for j := 2; j <= n; j++ {
			if i < j {
				dp[i][j] = dp[i][i]
			} else {
				dp[i][j] = dp[i-j][j] + dp[i][j-1]
			}
		}
	}

	return dp[m][n]
}

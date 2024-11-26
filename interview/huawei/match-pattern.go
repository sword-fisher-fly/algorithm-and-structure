package huawei

//https://www.nowcoder.com/practice/43072d50a6eb44d2a6c816a283b02036?tpId=37&tqId=21294&rp=1&ru=/exam/oj/ta&qru=/exam/oj/ta&sourceUrl=%2Fexam%2Foj%2Fta%3Fpage%3D2%26tpId%3D37%26type%3D37&difficulty=undefined&judgeStatus=undefined&tags=&title=
func MatchPattern(p string, pIdx int, t string, tIdx int) bool {
	if pIdx == len(p) && tIdx == len(t) {
		return true
	}

	if pIdx == len(p) || tIdx == len(t) {
		return false
	}

	if p[pIdx] == '?' {
		return MatchPattern(p, pIdx+1, t, tIdx+1)
	} else if p[pIdx] == '*' {
		return MatchPattern(p, pIdx+1, t, tIdx) || MatchPattern(p, pIdx+1, t, tIdx+1) || MatchPattern(p, pIdx, t, tIdx+1)
	} else if p[pIdx] == t[tIdx] {
		return MatchPattern(p, pIdx+1, t, tIdx+1)
	}

	return false
}

func MatchPatternByDy(p string, s string) bool {
	dp := make([][]bool, len(p)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s)+1)
	}

	dp[0][0] = true
	for i := 1; i <= len(p); i++ {
		if p[i-1] == '*' {
			dp[i][0] = true
		} else {
			break
		}
	}

	for i := 1; i <= len(p); i++ {
		for j := 1; j <= len(s); j++ {
			if p[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}

			if p[i-1] == '?' && (s[j-1] >= 'a' && s[j-1] <= 'z' || s[j-1] >= '0' && s[j-1] <= '9') {
				dp[i][j] = dp[i-1][j-1]
			} else if p[i-1] == '*' && (s[j-1] >= 'a' && s[j-1] <= 'z' || s[j-1] >= '0' && s[j-1] <= '9') {
				dp[i][j] = dp[i-1][j] || dp[i][j-1] || dp[i-1][j-1]
			}
		}
	}

	return dp[len(p)][len(s)]
}

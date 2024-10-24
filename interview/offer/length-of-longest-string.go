package offer

// TODO:

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/

func LengthOfLongestString(s string) int {
	leftIdx := 0

	ans := 0
	cIndex := [128]bool{}
	for i := range s {
		for cIndex[s[i]] {
			cIndex[s[leftIdx]] = false
			leftIdx++
		}

		ans = max(ans, i-leftIdx+1)
		cIndex[s[i]] = true
	}

	return ans
}

func LengthOfLongestSubstringII(s string) int {
	leftIdx := 0

	ans := 0
	sMap := make(map[byte]bool)
	for i := range s {
		for sMap[s[i]] {
			delete(sMap, s[leftIdx])
			leftIdx++
		}

		ans = max(ans, i-leftIdx+1)
		sMap[s[i]] = true
	}

	return ans
}

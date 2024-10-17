package twopointer

// https://leetcode.cn/problems/find-all-anagrams-in-a-string/
func FindAnagrams(s string, p string) []int {
	sl, pl := len(s), len(p)

	if sl < pl {
		return []int{}
	}

	ans := []int{}

	sCntMap := [26]int{}
	pCntMap := [26]int{}
	for i := range p {
		pCntMap[p[i]-'a']++
		sCntMap[s[i]-'a']++
	}

	isAnagram := func(p, q [26]int) bool {
		for i := 0; i < 26; i++ {
			if p[i] != q[i] {
				return false
			}
		}
		return true
	}

	if isAnagram(sCntMap, pCntMap) {
		ans = append(ans, 0)
	}

	for i := pl; i < sl; i++ {
		sCntMap[s[i]-'a']++
		sCntMap[s[i-pl]-'a']--
		if isAnagram(sCntMap, pCntMap) {
			ans = append(ans, i-pl+1)
		}
	}

	return ans
}

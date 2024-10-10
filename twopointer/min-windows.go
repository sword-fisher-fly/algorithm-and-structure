package twopointer

// 给定字符串s和目标字符串t
// 寻找在s中包含t所有字符的最小子串
// 假设: s和t都是由小写字母组成的字符串
// s: "wordgoodgoodgoodbestword"
// t: "goodword"
// Output: "wordgood"

func MinWindows(s, t string) string {
	res := ""
	left, right := 0, 0
	finalLeft, finalRight := -1, -1
	minW := len(s) + 1
	count := 0

	sFreq := [26]int{}
	tFreq := [26]int{}

	// 统计t中字符出现的频次
	for i := range t {
		tFreq[t[i]-'a']++
	}

	for left < len(s) {
		// if right+1 < len(s) && count < len(t) {
		// 	sFreq[s[right+1]-'a']++
		// 	if sFreq[s[right+1]-'a'] <= tFreq[s[right+1]-'a'] {
		// 		count++
		// 	}
		// 	right++
		// }
		if right < len(s) && count < len(t) {
			sFreq[s[right]-'a']++
			if sFreq[s[right]-'a'] <= tFreq[s[right]-'a'] {
				count++
			}
			right++
		} else {
			if right-left+1 < minW && count == len(t) {
				minW = right - left + 1
				finalLeft = left
				finalRight = right
			}

			if sFreq[s[left]-'a'] == tFreq[s[left]-'a'] {
				count--
			}
			sFreq[s[left]-'a']--
			left++
		}
	}

	if finalLeft != -1 {
		res = s[finalLeft:finalRight]
	}

	return res
}

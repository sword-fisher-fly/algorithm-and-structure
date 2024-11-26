package dynamic

// dp[i]: i长度的前缀s字符串可以由一个或多个字典中字符串组成
func WorkBreak(s string, words []string) bool {
	wordDict := make(map[string]bool)
	for _, word := range words {
		wordDict[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ { // 背包
		for j := 0; j < i; j++ { // 物品
			if dp[j] && wordDict[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func WordBreakByDFS(s string, words []string) bool {
	wordMap := make(map[string]bool)
	for i := range words {
		wordMap[words[i]] = true
	}

	umap := make(map[int]bool)
	var dfs func(int) bool
	dfs = func(startIndex int) bool {
		if startIndex == len(s) {
			return true
		}

		if umap[startIndex] {
			return false
		}

		for i := startIndex; i < len(s); i++ {
			if _, ok := wordMap[s[startIndex:i+1]]; ok && dfs(i+1) {
				return true
			}
		}

		umap[startIndex] = true
		return false
	}

	ret := dfs(0)
	return ret
}

func WordBreakByDFSII(s string, wordDict []string) bool {
	umap := make(map[int]bool)
	var dfs func(int) bool
	dfs = func(startIndex int) bool {
		if startIndex == len(s) {
			return true
		}

		if umap[startIndex] {
			return false
		}

		for _, k := range wordDict {
			if startIndex+len(k) > len(s) {
				continue
			}

			str := s[startIndex : startIndex+len(k)]
			if str == k && dfs(startIndex+len(k)) {
				return true
			}
		}

		umap[startIndex] = true

		return false
	}

	return dfs(0)
}

// ?? Where goes wrong？
// func WorkBreakII(s string, words []string) bool {
// 	wordDict := make(map[string]bool)
// 	for _, word := range words {
// 		wordDict[word] = true
// 	}

// 	dp := make([]bool, len(s)+1)
// 	dp[0] = true

// 	for i := 0; i < len(words); i++ { // 物品
// 		for j := len(words[i]); j <= len(s); j++ { // 背包
// 			word := s[j-len(words[i]) : j-len(words[i])+len(words[i])]

// 			fmt.Printf("word=%s\n", word)
// 			if dp[j-len(words[i])] && word == words[i] {
// 				dp[j] = true
// 			}
// 		}
// 	}

// 	return dp[len(s)]
// }
